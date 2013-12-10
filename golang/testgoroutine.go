package main

import (
	"time"
)

func xxx() {
	for {
		time.Sleep(1 * time.Second)
		panic("XXX")
	}
}

func main() {
	go xxx()

	time.Sleep(10 * time.Second)
}
