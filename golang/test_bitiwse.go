package main

import (
	"fmt"
)

func main() {
	x := uint32(128)
	id := int32(128)

	y := uint64(id)<<32 | uint64(x)
	fmt.Println(y)
}
