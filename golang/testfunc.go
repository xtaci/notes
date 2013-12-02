package main

import (
	"time"
	"fmt"
)

func test(a, b int) int {
	return a+b
}

func main() {
	start := time.Now()
	test(1,1)
	end :=  time.Now()
	fmt.Println(end.Sub(start))
}
