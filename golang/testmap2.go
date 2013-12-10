package main

import (
	"fmt"
)

func main() {
	m := make(map[int32]chan int)

	select {
	case m[2] <- 1:
	default:
		fmt.Println("error")
	}
}
