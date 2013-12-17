package main

import (
	"fmt"
)

type X struct {
	A int
	B string
	X []int
}

func main() {
	x := X{}
	fmt.Println(cap(x.X), len(x.X))
	x.X = append(x.X[:0], x.X[1:]...)
}
