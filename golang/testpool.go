package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{}
	fmt.Println(pool.Get())

	m := make([]int, 0, 10)
	fmt.Println(len(m[:10]))
	fmt.Println(m.([]int))
}
