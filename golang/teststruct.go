package main

import (
	"fmt"
)

type X struct {
	A int
	B string
}

func main(){
	x := X{1, "hello"}
	fmt.Println(x)
}
