package main

import (
	"fmt"
	"time"
)

func main() {
	const N = 1000000
	x := make([]int32, N)
	x[0] = 10

	start := time.Now()
	y := append(x[1:], x[0])
	x = y
	fmt.Println(time.Now().Sub(start))
}
