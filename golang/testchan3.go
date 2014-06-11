package main

import (
	"time"
)

func main() {
	//var x chan int
	//
	x := make(chan int)
	close(x)
	go func() { x <- 1 }()
	go func() { println(<-x) }()
	time.Sleep(time.Hour)
}
