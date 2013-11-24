package main

import (
	"reflect"
	"strconv"
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

	f, err := strconv.ParseFloat("0", 32)
	fmt.Println(f, err)

	f1 := float32(0)
	f2 := float32(0)
	fmt.Println(f1/f2 < 0)

	testx := float32(90000000000)
	fmt.Println(testx)

	d1, err := strconv.ParseFloat("0.2", 32)
	d2, err := strconv.ParseFloat("0.4", 32)
	fmt.Println(d1,d2,d1+d2)
}
