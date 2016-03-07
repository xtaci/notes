package main

import (
	"fmt"
)

func main() {
	later := uint32(0)
	early := uint32(1)
	fmt.Println(later - early)
	fmt.Println(int32(later - early))
}
