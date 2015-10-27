package main

import (
	"fmt"
)

func main() {
	x := [3]int{10, 9, 8}
	for k := range x {
		fmt.Println(k)
	}
}
