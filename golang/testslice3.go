package main

import (
	"fmt"
)

func main() {
	xx := []int32{10, 20, 30, 40, 12, 3, 4, 5, 5}
	a := xx[:2]
	b := xx[:3]
	fmt.Println(cap(a), cap(b))
	fmt.Println(a[:])
}
