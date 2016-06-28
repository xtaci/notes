package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(-1 % 8)
	start := time.Now()
	i := uint32(1)
	for {
		if i == 0 {
			break
		}
		i++
	}
	fmt.Println(time.Now().Sub(start))

}
