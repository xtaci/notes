package main

import (
	"time"
)

func main() {
	x := make(chan int)
	close(x)

	x = nil
	select {
	case x <- 1:
		println("x")
	case <-time.After(time.Second):
		println("time")
	}
}
