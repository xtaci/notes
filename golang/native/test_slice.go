package main

import "fmt"

func main() {
	x := make([]byte, 10)
	y := x[:5]
	fmt.Println(cap(y))
	y = y[0:1:2]
	fmt.Println(cap(y))
}
