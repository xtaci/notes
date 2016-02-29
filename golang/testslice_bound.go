package main

import "fmt"

func main() {
	x := []int{1}
	fmt.Println("len:", len(x), "cap:", cap(x))
	fmt.Println(x[len(x):len(x)])
}
