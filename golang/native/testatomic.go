package main

import "sync/atomic"

func main() {
	var x uint64
	x = 100
	atomic.AddUint64(&x, 1)
	println(x)
	atomic.AddUint64(&x, ^uint64(0))
	println(x)
}
