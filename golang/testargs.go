package main

import (
	"fmt"
)

type X struct {
	x int32
}

func margs(args ...X) {
	for k := range args {
		fmt.Println(args[k])
	}
}

func main() {
	x1 := X{1}
	x2 := X{2}
	margs(x2, x2)
}
