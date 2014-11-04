package main

import (
	"fmt"
	"time"
)

import (
	"utils"
)

func main() {
	const N = 100000
	m := make(map[int]int)
	for i := 0; i < N; i++ {
		m[i] = i
	}

	start := time.Now()
	r := <-utils.LCG % uint32(N)
	for _ = range m {
		if r == 0 {
			break
		}
		r--
	}

	fmt.Println(time.Now().Sub(start))
}
