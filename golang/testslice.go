package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0,0)
	fmt.Println(append(s[:0], s[1:]...))
}
