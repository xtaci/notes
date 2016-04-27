package main

import "fmt"

type test struct {
	A int
}

func main() {
	cnt := 0
L:
	n := test{cnt}
	defer fmt.Println(n)

	if cnt > 1 {
		return
	}
	cnt++
	goto L
}
