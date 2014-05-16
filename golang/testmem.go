package main

import (
	"fmt"
	"time"
)

const SIZE = 1024 * 1024
const COUNT = 4096

func main() {
	start := time.Now()
	for k := 0; k < COUNT; k++ {
		sl := make([]byte, SIZE)
		for i := 0; i < SIZE; i++ {
			sl[i] = 1
		}
	}

	fmt.Println(time.Now().Sub(start))
}
