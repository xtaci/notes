package main

import (
	"fmt"
	"time"
)

func main() {
	const N = 10000000
	s := make([]int, N)
	start := time.Now()
	for i := 0; i < N; i++ {
		s = s[1:]
		s = append(s, 1)
	}

	fmt.Println(time.Now().Sub(start))
}
