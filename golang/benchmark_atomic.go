package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var x int64
	start := time.Now()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&x, 1)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
