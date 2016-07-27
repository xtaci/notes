package main

import "fmt"

const y = 10

func main() {
	const x = &y
	fmt.Println(x)
}
