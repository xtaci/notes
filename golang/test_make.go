package main

import "fmt"

func main() {
	x := make([]int, 0)
	fmt.Println(len(x), cap(x))
}
