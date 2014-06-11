package main

import (
	"time"
)

var x0 uint32 = uint32(time.Now().UnixNano())
var a uint32 = 1664525
var c uint32 = 1013904223

var LCG chan uint32

const PRERNG = 128

//------------------------------------------------ 全局快速随机数发生器
func init() {
	LCG = make(chan uint32, PRERNG)
	go func() {
		for {
			x0 = a*x0 + c
			LCG <- x0
		}
	}()
}

func main() {
	a := <-LCG
	println(a)
	cnt := 0
	for {
		b := <-LCG
		cnt++
		if b == a {
			break
		}
		if cnt%1000000 == 0 {
			println(cnt, b)
		}
	}
	println("cycle:", cnt)
}
