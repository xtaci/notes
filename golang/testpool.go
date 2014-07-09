package main

import (
	"fmt"
	"sync"
)

func main() {
	p := sync.Pool{}
	p.New = func() interface{} { return make([]int, 10) }
	fmt.Println(p.Get())
}
