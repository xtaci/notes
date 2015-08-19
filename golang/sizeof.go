package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := struct{ _ int }{}
	y := []bool{}
	z := map[int]struct{}{}
	a := make(chan struct{})
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Sizeof(y))
	fmt.Println(unsafe.Sizeof(z))
	fmt.Println(unsafe.Sizeof(a))
}
