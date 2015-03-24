package main

import (
	"fmt"
)

func ret2() (string, string) {
	return "a", "b"
}

func ret3() (string, string, string) {
	return "a", "b", "c"
}

func main() {
	fmt.Println("ret2:", ret2())
	fmt.Println("ret3:", ret3())
}
