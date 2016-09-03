package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := make([]byte, 1024)
	fmt.Println(unsafe.Pointer(&arr[0]))
	arr = append(arr, make([]byte, 8192)...)
	fmt.Println(unsafe.Pointer(&arr[0]))
}
