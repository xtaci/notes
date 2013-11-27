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
	s[0] = 10
	fmt.Println(<-c)

	close(c)
	fmt.Println(<-c)
}
