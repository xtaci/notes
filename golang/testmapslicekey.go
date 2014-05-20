package main

import (
	"fmt"
)

func main() {
	data := make(map[[]byte]string)

	s1 := []byte{1, 2, 3}
	s2 := []byte{1, 2, 3}

	data[s1] = "a"
	data[s2] = "b"

	fmt.Println(data)
}
