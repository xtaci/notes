package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var arr [4]int
	a := uintptr(unsafe.Pointer(&arr))
	var s1 = struct {
		addr uintptr
		len  int
		cap  int
	}{a, 4, 4}

	s := *(*[]int)(unsafe.Pointer(&s1))
	for k := range s {
		fmt.Println(k, s[k])
		s[k] = k
	}

	fmt.Println(s)
}
