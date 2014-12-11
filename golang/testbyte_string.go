package main

import (
	"fmt"
)

func main() {
	var x []byte
	if string(x) == "" {
		fmt.Println(true)
	}
}
