package main

import (
	"fmt"
	"time"
)

func main() {
	a := 1

	go func() {
		a = 10

		fmt.Printf("x: %p\n", &a)
	}()

	for _ = range time.Tick(time.Second) {
		fmt.Printf("x: %p,%d\n", &a, a)
	}

}
