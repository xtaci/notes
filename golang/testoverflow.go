package main

import (
	"fmt"
	"time"
)

var i = uint32(1)

func main() {
	stack()
	heap()
}

func stack() {
	start := time.Now()
	i := uint32(1)
	for i != 0 {
		i++
	}
	fmt.Println("stack:", time.Now().Sub(start))
}

func heap() {
	start := time.Now()
	for i != 0 {
		i++
	}
	fmt.Println("heap:", time.Now().Sub(start))
}
