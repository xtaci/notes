package main

import (
	"fmt"
	"log"
	"net"
	"sync/atomic"
	"time"
)

const (
	srvAddr         = "224.0.0.1:9999"
	maxDatagramSize = 8192
)

var (
	count uint64
)

func main() {
	for i := 0; i < 5000; i++ {
		go serveMulticastUDP(srvAddr, msgHandler)
	}
	go ping(srvAddr)
	select {}
}

func ping(a string) {
	addr, err := net.ResolveUDPAddr("udp", a)
	if err != nil {
		log.Fatal(err)
	}
	c, err := net.DialUDP("udp", nil, addr)
	n := 0
	for {
		c.Write([]byte(fmt.Sprintf("hello, world%v\n", n)))
		time.Sleep(10 * time.Millisecond)
		n++
	}
}

func msgHandler(src *net.UDPAddr, n int, b []byte) {
	cnt := atomic.AddUint64(&count, 1)
	if cnt%5000 == 0 {
		println(count)
	}
}

func serveMulticastUDP(a string, h func(*net.UDPAddr, int, []byte)) {
	addr, err := net.ResolveUDPAddr("udp", a)
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.ListenMulticastUDP("udp", nil, addr)
	l.SetReadBuffer(maxDatagramSize)
	for {
		b := make([]byte, maxDatagramSize)
		n, src, err := l.ReadFromUDP(b)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		h(src, n, b)
	}
}
