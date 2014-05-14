package main

import (
	"fmt"
	"time"
)

type X struct {
	key   string
	value int
}

const N = 10000000

func main() {
	test := make(map[string]int)

	for k := 0; k < N; k++ {
		test[fmt.Sprint(k)] = k
	}

	fmt.Println(test)
	time.Sleep(time.Minute)
}
