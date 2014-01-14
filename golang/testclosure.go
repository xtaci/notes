package main

import (
	"fmt"
	"time"
)

func main() {
	x := 100
	send := func(a int) {
		fmt.Println(a, x)
	}

	for i := 0; i < 10; i++ {
		go send(i)
		send(i)
	}
	time.Sleep(time.Second)
}
