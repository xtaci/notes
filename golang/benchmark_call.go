package main

import (
	"fmt"
	"misc/packet"
	"time"
)

type Packet struct {
	pos  uint
	data []byte
}

var _dum []byte

func init() {
	_dum = make([]byte, 10)
}

func get() []byte {
	return _dum
}

func main() {
	var ret []byte
	<-time.After(2 * time.Second)
	start := time.Now()
	N := 1
	for i := 0; i < N; i++ {
		ret = packet.Pack(1, nil, nil)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println(ret)
}
