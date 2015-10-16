package main

import "fmt"

func main() {
	for _, s := range []string{"foo", "bar"} {
		x := s
		func(s string) {
			fmt.Printf("s: %s\n", s)
			fmt.Printf("x: %s\n", x)
		}(s)
	}
}
