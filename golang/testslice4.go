package main

import (
	"fmt"
)

func main() {
	x := make([]int, 2)
	x[1] = 1

	fmt.Println(x[1:1])
	y := x[:100]
	fmt.Println(len(y))
}
