package main

import "fmt"

func main() {
	defer re()
	f()
}
func re() {
	fmt.Println(recover())
}

func f() {
	panic("test")
}
