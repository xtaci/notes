package main

import (
	"fmt"
	"misc/packet"
	"time"
)

var (
	_dum []byte
	pool chan *packet.Packet
)

func init() {
	_dum = make([]byte, 10)

	pool = make(chan *packet.Packet, 1024)
	go func() {
		for {
			pool <- packet.Writer()
		}
	}()
}

func get() []byte {
	return _dum
}

func main() {
	var ret interface{}
	<-time.After(2 * time.Second)
	start := time.Now()
	N := 1
	for i := 0; i < N; i++ {
		//ret = packet.Pack(1, nil, &packet.Packet{})
		//ret = packet.Writer()
		ret = <-pool
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println(ret)
}
