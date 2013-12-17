package main

import "reflect"
import "fmt"

func main() {
	x := []byte{1, 2}
	y := "abc"
	fmt.Printf("%T\n", x)
	fmt.Println(reflect.TypeOf(x).String())
	fmt.Println(reflect.TypeOf(y).String())
}
