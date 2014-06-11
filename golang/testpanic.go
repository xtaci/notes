package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		func() {
			defer func() {
				if x := recover(); x != nil {
					fmt.Println(x)
				}
			}()
			panic("test")
		}()
	}
}
