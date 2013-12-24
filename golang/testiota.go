package main

import (
	"fmt"
)

const (
	STAR_0 = int8(iota)
	STAR_1
	STAR_2
	STAR_3

	STAR_100 = 100
)

func main() {
	fmt.Println(STAR_0)
	fmt.Println(STAR_3)
	fmt.Println(STAR_100)
}
