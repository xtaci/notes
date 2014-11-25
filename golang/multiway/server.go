package main

import (
	"encoding/binary"
	"hash/crc32"
	"log"
	"net"
	"os"
	"time"
)

import (
	"utils"
)

const (
	KB           = 1 << 10
	UDP_BUF_SIZE = 64 * KB // 最大的上报的包大小64K
)

//----------------------------------------------- 启动多路UDP接收器
func start_multiway_server() {
	for k := range WAYS {
		addr, err := net.ResolveUDPAddr("udp", WAYS[k])
		checkError(err)

		conn, err := net.ListenUDP("udp", addr)
		checkError(err)

		utils.INFO("MULTIWAY SERVER ON UDP:", WAYS[k])
		go handle_server(conn)
	}

	for {
		<-time.After(time.Second)
	}
}

//----------------------------------------------- 作为服务器端的UDP处理
func handle_server(conn *net.UDPConn) {
	for {
		buf := make([]byte, FRAGMENT_SIZE*2)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			utils.ERR("read udp failed", n, addr, err)
			continue
		}

		if n < 12 {
			continue
		}

		crc := binary.BigEndian.Uint32(buf[:4])
		streamid := binary.BigEndian.Uint32(buf[4:8])
		seqid := binary.BigEndian.Uint32(buf[8:12])
		if crc != crc32.ChecksumIEEE(buf[12:n]) {
			continue
		}
		utils.DEBUG("handle_server:", n, streamid, seqid)
		ServerStreamInstance.recv(streamid, seqid, buf[12:n])
	}
}

func checkError(err error) {
	if err != nil {
		log.Println("Fatal error:", err)
		os.Exit(-1)
	}
}
