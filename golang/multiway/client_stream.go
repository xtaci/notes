package main

import (
	"hash/crc32"
	"net"
	"sync"
	"time"
)

import (
	"misc/packet"
	"utils"
)

var (
	WAYS = []string{"106.187.101.141:60000", "106.187.101.141:60001", "106.187.101.141:60002"}
)

type ClientStream struct {
	max_confirmed uint32            // 当前成功投递到TCP的包
	max_seqid     uint32            // 收到的最大包序号
	wind          map[uint32][]byte // 接收窗口
	conn          net.Conn          // 对应到GS/或Client的连接
}

type ClientStreamManager struct {
	all     map[uint32]*ClientStream // 所有连接
	ways    []*net.UDPConn
	pending chan []byte
	delayed chan []byte
	sync.Mutex
}

var (
	ClientStreamInstance ClientStreamManager
)

func init() {
	ClientStreamInstance.init()
	go ClientStreamInstance.sender()
	go ClientStreamInstance.delayed_sender()
}

//---------------------------------------------------------- 初始化客户断流
func (sm *ClientStreamManager) init() {
	sm.all = make(map[uint32]*ClientStream)
	sm.ways = make([]*net.UDPConn, 0)
	sm.pending = make(chan []byte, 65535)
	sm.delayed = make(chan []byte, 65535)
	for k := range WAYS {
		addr, err := net.ResolveUDPAddr("udp", WAYS[k])
		checkError(err)
		conn, err := net.DialUDP("udp", nil, addr)
		checkError(err)
		sm.ways = append(sm.ways, conn)
	}
}

//---------------------------------------------------------- 从server收取UDP包发送到client
func (sm *ClientStreamManager) connect(streamid uint32, conn net.Conn) {
	sm.Lock()
	defer sm.Unlock()
	stream := &ClientStream{max_confirmed: 0,
		max_seqid: 0,
		wind:      make(map[uint32][]byte),
		conn:      conn,
	}
	sm.all[streamid] = stream
}

func (sm *ClientStreamManager) close(streamid uint32) {
	sm.Lock()
	defer sm.Unlock()
	delete(sm.all, streamid)
}

//---------------------------------------------------------- 从server收取UDP包发送到client
func (sm *ClientStreamManager) recv(streamid uint32, seqid uint32, payload []byte) {
	sm.Lock()
	defer sm.Unlock()

	stream, ok := sm.all[streamid]
	if !ok {
		return
	}

	if stream.max_confirmed >= seqid { // 丢弃重复包
		return
	}

	// 新的seqid出现，尝试发送
	stream.wind[seqid] = payload
	if seqid > stream.max_seqid {
		stream.max_seqid = seqid
	}
	utils.DEBUG("max confirmed", stream.max_confirmed)
	utils.DEBUG("max seqid", stream.max_seqid)

	for stream.max_confirmed < stream.max_seqid {
		payload, ok := stream.wind[stream.max_confirmed+1]
		utils.DEBUG("recv seqid", stream.max_confirmed+1)
		if ok { // 发送
			stream.conn.Write(payload)
			stream.max_confirmed++
			// 清空发送过的数据包
			delete(stream.wind, stream.max_confirmed)
		} else {
			break
		}
	}
}

//---------------------------------------------------------- 客户端发送(TCP->UDP)
func (sm *ClientStreamManager) send(streamid uint32, seqid uint32, data []byte) {
	writer := packet.Writer()
	writer.WriteU32(crc32.ChecksumIEEE(data))
	writer.WriteU32(streamid)
	writer.WriteU32(seqid)
	writer.WriteRawBytes(data)
	sm.pending <- writer.Data()
}

func (sm *ClientStreamManager) sender() {
	for {
		data := <-sm.pending
		for k := range sm.ways {
			sm.ways[k].Write(data)
		}

		sm.delayed <- data
	}
}

func (sm *ClientStreamManager) delayed_sender() {
	for {
		<-time.After(time.Second)
		sz := len(sm.delayed)
		for i := 0; i < sz; i++ {
			data := <-sm.delayed
			for k := range sm.ways {
				sm.ways[k].Write(data)
			}
		}
	}
}
