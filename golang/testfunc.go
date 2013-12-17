package main

import (
	"fmt"
	"time"
)

func test(a, b int) int {
	return a + b
}

func main() {
	time.Sleep(time.Second)
	//	x := test
	start := time.Now()
	//	x(1,1)
	test(1, 1)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
