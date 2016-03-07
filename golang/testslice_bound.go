package main

import "fmt"

func main() {
	x := []int{1}
	fmt.Println("len:", len(x), "cap:", cap(x))
	fmt.Println(x[len(x):len(x)])

	var y []int
	y = []int{1}
	z := y[1:0]
	fmt.Println(len(z), cap(z))
}
