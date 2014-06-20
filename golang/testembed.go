package main

import "fmt"

type T interface {
	Print()
}

type X struct {
	v []int
}

func (x X) Print() {
	fmt.Println("X print", x.v)
}

type X2 struct {
}

func (x X2) Print() {
	fmt.Println("X2 print")
}

type Y struct {
	X
	X2
}

func main() {
	x := X{[]int{10, 20}}
	x.Print()
	Y{x, X2{}}.Print()
}
