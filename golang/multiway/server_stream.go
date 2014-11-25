package main

import (
	"encoding/binary"
	"hash/crc32"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

import (
	"misc/packet"
	"utils"
)

var (
	FRONTS = []string{"112.124.117.14:50000", "112.124.117.14:50001", "112.124.117.14:50002"}
)

type ServerStream struct {
	max_confirmed uint32            // 当前成功投递到TCP的包
	max_seqid     uint32            // 收到的最大包序号
	wind          map[uint32][]byte // 接收窗口
	conn          net.Conn          // 对应到GS/或Client的连接
}

type ServerStreamManager struct {
	all     map[uint32]*ServerStream // 所有连接
	fronts  []*net.UDPConn
	pending chan []byte
	delayed chan []byte
	sync.Mutex
}

var (
	ServerStreamInstance ServerStreamManager
)

func init() {
	ServerStreamInstance.init()
	go ServerStreamInstance.sender()
	go ServerStreamInstance.delayed_sender()
}

//---------------------------------------------------------- 初始化服务器流
func (sm *ServerStreamManager) init() {
	sm.all = make(map[uint32]*ServerStream)
	sm.fronts = make([]*net.UDPConn, 0)
	sm.pending = make(chan []byte, 65535)
	sm.delayed = make(chan []byte, 65535)
	for k := range FRONTS {
		addr, err := net.ResolveUDPAddr("udp", FRONTS[k])
		checkError(err)
		conn, err := net.DialUDP("udp", nil, addr)
		checkError(err)
		sm.fronts = append(sm.fronts, conn)
	}
}

//---------------------------------------------------------- 从front收取UDP发送到GS
func (sm *ServerStreamManager) recv(streamid uint32, seqid uint32, payload []byte) {
	sm.Lock()
	defer sm.Unlock()

	utils.DEBUG("stream send", streamid, seqid)

	stream, ok := sm.all[streamid]
	if !ok { // 开启新连接
		c, err := net.Dial("tcp", "54.201.148.188:8888")
		if err != nil {
			// handle error
			log.Println(err)
		}

		log.Println("open connection")

		stream = &ServerStream{max_confirmed: 0,
			max_seqid: 0,
			wind:      make(map[uint32][]byte),
			conn:      c,
		}
		sm.all[streamid] = stream
		go sm.receiver(streamid, c)
	}

	if stream.max_confirmed >= seqid { // 丢弃重复包
		return
	}

	// 新的seqid出现，尝试发送
	stream.wind[seqid] = payload
	// 更新最大的seqid
	if seqid > stream.max_seqid {
		stream.max_seqid = seqid
	}
	utils.DEBUG("max confirmed", stream.max_confirmed)
	utils.DEBUG("max received", stream.max_seqid)

	for stream.max_confirmed < stream.max_seqid {
		payload, ok := stream.wind[stream.max_confirmed+1]
		utils.DEBUG("send seqid", stream.max_confirmed+1)
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

func (sm *ServerStreamManager) receiver(streamid uint32, conn net.Conn) {
	// network loop
	seqid := uint32(0)
	defer func() {
		sm.Lock()
		defer sm.Unlock()
		delete(sm.all, streamid)
	}()

	for {
		header := make([]byte, 2)
		n, err := io.ReadFull(conn, header)
		if err != nil {
			utils.WARN("error receiving header, bytes:", n, "reason:", err)
			return
		}
		size := binary.BigEndian.Uint16(header)

		// read msg
		data := make([]byte, size)
		n, err = io.ReadFull(conn, data)
		if err != nil {
			utils.WARN("error receiving msg, bytes:", n, "reason:", err)
			return
		}

		// 完整packet
		data = append(header, data...)
		// 分片
		fragments := fragment(data)
		for k := range fragments {
			seqid++
			writer := packet.Writer()
			writer.WriteU32(crc32.ChecksumIEEE(fragments[k]))
			writer.WriteU32(streamid)
			writer.WriteU32(seqid)
			writer.WriteRawBytes(fragments[k])
			sm.pending <- writer.Data()
		}
	}
}

func (sm *ServerStreamManager) sender() {
	for {
		data := <-sm.pending
		for k := range sm.fronts {
			sm.fronts[k].Write(data)
		}
		sm.delayed <- data
	}
}

func (sm *ServerStreamManager) delayed_sender() {
	for {
		<-time.After(time.Second)
		sz := len(sm.delayed)
		for i := 0; i < sz; i++ {
			data := <-sm.delayed
			for k := range sm.fronts {
				sm.fronts[k].Write(data)
			}
		}
	}
}
