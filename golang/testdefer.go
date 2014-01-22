package main

import (
	"fmt"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	defer func() {
		fmt.Println("###P2")
		panic("P2")
		fmt.Println("###P3")
	}()

	panic("xx")
}
