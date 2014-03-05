package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	for i := 0; i < 1000000; i++ {
		time.Now()
	}
	fmt.Println(time.Now())
}
