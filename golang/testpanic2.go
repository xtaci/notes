package main

import (
	"fmt"
)

func main() {
	defer func() { println("A") }()
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	defer func() { println("B") }()

	panic("XXXX")
}
