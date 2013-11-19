package main

import (
	"reflect"
	"fmt"
)
func main() {
	x := int64(100000)
	const Y = 3*60
	// comparable 
	if (x > Y*1.1) {
		fmt.Println(reflect.TypeOf(Y*1.1))
		fmt.Println(reflect.TypeOf(x))
	}
}
