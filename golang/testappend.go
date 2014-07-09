package main

import (
	"fmt"
)

func main() {
	x := make([]byte, 2)
	var y []byte
	x = append(x, y...)
	fmt.Println(x)
}
