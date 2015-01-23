package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10)
	fmt.Println(len(x[10:]))
}
