package main

import (
	"fmt"
)

func main() {
	fmt.Println(1<<0)
	fmt.Println(1<<1)
	fmt.Println((uint32(1) << 0) -1)
}
