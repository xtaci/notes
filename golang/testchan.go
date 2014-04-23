package main

import (
	"fmt"
)

func main() {
	// pass a slice to chan is by reference
	c := make(chan []byte, 10)
	s := make([]byte, 1)
	s[0] = 0
	c <- s[:]
	// modify orginal data
	s[0] = 10

	// read from channel
	fmt.Println(<-c)

	close(c)
	doubleclose()
}

func doubleclose() {
	c := make(chan []byte, 10)
	close(c)
	close(c)
}
