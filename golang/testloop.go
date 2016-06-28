package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		x := false
		if i == 5 {
			x = true
		}

		fmt.Println(x)
	}
}
