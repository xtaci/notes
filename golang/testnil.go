package main

import (
	"fmt"
)

func isnil(x interface{}) bool {
	if x == nil {
		return true
	}

	return false
}

func main() {
	var x1 chan byte
	var x2 map[int]int
	fmt.Println(isnil(x1))
	fmt.Println(isnil(x2))
	fmt.Println(isnil(x2))
}
