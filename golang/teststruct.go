package main

import (
	"fmt"
)

type X struct {
	A int
	B string
	X []int
}

func (x *X) Fun1() {
	x.A = 1
}

type AliasX X

func (ax *AliasX) Fun2() {
	ax.B = "1"
}

func main() {
	x := &X{}
	x.Fun1()
	x.Fun2()
}
