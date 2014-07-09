package main

import (
	"fmt"
	"time"
)

type X struct {
	A int
}

func main() {
	send := func(x X) {
		fmt.Println(x)
	}

	for i := 0; i < 10; i++ {
		x := X{i}
		go send(x)
	}
	time.Sleep(time.Second)
}
