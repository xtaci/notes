package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	x := make(map[int]int)
	fmt.Println(xml.Marshal(x))
}
