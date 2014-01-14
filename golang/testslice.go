package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	slice := arr[1:2]
	slice = append(slice, 6, 7, 8)
	fmt.Println(slice)
	fmt.Println(arr)
}
