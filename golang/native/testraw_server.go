package main

import (
	"fmt"
	"net"
)

func main() {
	laddr, _ := net.ResolveIPAddr("ip4", "0.0.0.0:1234")
	conn, err := net.ListenIP("ip4:tcp", laddr)
	if err != nil {
		panic(err)
	}
	for {
		buf := make([]byte, 1024)
		numRead, raddr, err := conn.ReadFrom(buf)
		fmt.Println(buf, raddr, numRead, err)
	}
}
