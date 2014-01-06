package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)
	v, ok := m[""]
	fmt.Println(v, ok)
	m[""] = ""
	v, ok = m[""]
	fmt.Println(v, ok)
	var x map[string]string
	v, ok = x["xx"]
	fmt.Println(v, ok)
}
