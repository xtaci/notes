package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		x := new(int)
		go func() {
			fmt.Printf("%p %v \n", x, *x)
		}()
	}
	time.Sleep(time.Second)
}
