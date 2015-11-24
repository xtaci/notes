package main

import (
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		x := i
		f := func() {
			println(i, x)
		}
		go f()
	}
	time.Sleep(time.Minute)
}
