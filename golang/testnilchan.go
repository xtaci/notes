package main

import (
	"time"
)

func main() {
	var c1 chan byte
	c2 := make(chan byte)
	close(c2)
	go func() {
		for {
			select {
			case x, ok := <-c1:
				println("c1", x, ok)
			case c2<-1:
				println("c2")
			}
		}
	}()

	time.Sleep(time.Minute)
}
