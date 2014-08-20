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

func main() {
	start := time.Now()
	N := 1
	for i := 0; i < N; i++ {
		packet.Pack(1, nil)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
