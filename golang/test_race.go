package main

import (
	"fmt"
	//"time"
)

var x int

func main() {
	ch := make(chan bool, 1)
	go func() {
		<-ch
		x = 100
	}()

	ch <- true
	for i := 0; i < 10; i++ {
		fmt.Println(x)
		for j := 0; j < 100000000; j++ {
		}
	}
}
