package main

import (
	"fmt"
)

type X struct {
	A int
}

func main() {
	xs := make([]X, 10)
	xs[5] = X{5}
	fmt.Println(xs)
	y := &xs[5]

	y.A = 10
	fmt.Println(xs)
	fmt.Println(y)
	xs = xs[:1]
	fmt.Println(y)
}
