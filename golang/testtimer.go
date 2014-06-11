package main

import (
	"time"
)

func main() {
	ch := time.After(time.Second)

	x := 0
	for {
		select {
		case <-ch:
			x++
			println(x)
			ch = time.After(time.Second)
		}
	}
}
