package main

import (
	"fmt"
)

type X struct {
	A int32
}

func test(x interface{}) {
	m := x.(X)
	m.A = 5
}

func main() {
	x := X{1}
	test(x)
	fmt.Println(x)
}
