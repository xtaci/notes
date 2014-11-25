package main

import (
	"encoding/binary"
	"hash/crc32"
	"io"
	"net"
	"sync/atomic"
)

import (
	"utils"
)

var (
	_streamid uint32
)

var (
	_ways []*net.UDPConn
)

func start_multiway_front() {
	// start listening
	service := ":3000"
	utils.INFO("Service:", service)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	utils.INFO("Front Ready!!!")

	// start udp
	for k := range FRONTS {
		addr, err := net.ResolveUDPAddr("udp", FRONTS[k])
		checkError(err)

		conn, err := net.ListenUDP("udp", addr)
		checkError(err)

		utils.INFO("MULTIWAY CLIENT ON UDP:", FRONTS[k])
		go handle_client_udp(conn)
	}

	// loop accepting
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			utils.WARN("accept failed", err)
			continue
		}
		go handle_client_tcp(conn)
	}
}

//----------------------------------------------- start a goroutine when a new connection is accepted
func handle_client_tcp(conn *net.TCPConn) {
	// network loop
	streamid := atomic.AddUint32(&_streamid, 1)
	seqid := uint32(0)
	ClientStreamInstance.connect(streamid, conn)

	defer func() {
		ClientStreamInstance.close(streamid)
	}()

	for {
		header := make([]byte, 2)
		n, err := io.ReadFull(conn, header)
		if err != nil {
			utils.WARN("error receiving header, bytes:", n, "reason:", err)
			return
		}
		size := binary.BigEndian.Uint16(header)
		data := make([]byte, size)
		n, err = io.ReadFull(conn, data)

		if err != nil {
			utils.WARN("error receiving msg, bytes:", n, "reason:", err)
			return
		}

		// 完整packet
		data = append(header, data...)
		fragments := fragment(data)
		for k := range fragments {
			seqid++
			ClientStreamInstance.send(streamid, seqid, fragments[k])
		}
	}
}

//----------------------------------------------- 作为客户端的UDP
func handle_client_udp(conn *net.UDPConn) {
	for {
		// UDP接收缓冲区
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
		utils.DEBUG("handle_client_udp", n, streamid, seqid)
		ClientStreamInstance.recv(streamid, seqid, buf[12:n])
	}
}
