package main

import (
	"fmt"
)

type X struct {
	Data map[string]string
}

var tmp map[string]interface{}

func main() {
	tmp = make(map[string]interface{})
	x := &X{}
	x.Data["A"] = "B"
	fmt.Println(x)

	tmp["a"] = 1
	tmp["b"] = 2

	fmt.Println(tmp["a"] + tmp["b"])
}
