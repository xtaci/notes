package main

import (
	"fmt"
)

type X struct {
	Data map[string]string
}

func main() {
	x := &X{}
	x.Data["A"] = "B"
	fmt.Println(x)
}
