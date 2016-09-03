package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(f())
}

func f() (err error) {
	err = errors.New("cd")
	fmt.Printf("%v\n", &err)
	x, err := 1, 2
	fmt.Printf("%v\n", &err)
	fmt.Println(x)
	if err != nil {
		return err
	}
	return nil
}
