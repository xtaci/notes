package main

import (
	"fmt"
)

type X struct {
	A int
}

func main() {
	mm := make([]X, 10)

	temp := X{100}

	for k := range mm {
		mm[k] = temp
	}

	fmt.Println(mm)
}
