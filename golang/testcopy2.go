package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2}
	copy(x, x[1:])
	fmt.Println(x)
}
