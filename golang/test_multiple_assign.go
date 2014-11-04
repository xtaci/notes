package main

import (
	"fmt"
)

func main() {
	ids, cups := make([]int32, 10), make([]int32, 10)
	for i := int32(0); i < 10; i++ {
		ids[i] = 10 - i
		cups[i] = i
	}

	fmt.Println(ids, cups)
}
