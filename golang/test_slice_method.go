package main

import (
	"fmt"
)

type X struct {
	v int
}

func (x *X) set(v int) {
	x.v = v
}

func main() {
	ss := make([]X, 10)
	for i := 0; i < 10; i++ {
		ss[i].set(i)
	}

	fmt.Println(ss)
}
