package main

import (
	"time"
)

func main() {
	ch := time.NewTimer(time.Second)

	x := 0
	for {
		select {
		case <-ch.C:
			x++
			println(x)
			ch.Reset(time.Second)
		}
	}
}
