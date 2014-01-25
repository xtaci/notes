package main

import (
	"encoding/json"
	"fmt"
	"types/gifts"
)

func main() {
	x := []byte{123, 34, 73, 100, 34, 58, 48, 44, 34, 67, 111, 100, 101, 34, 58, 34, 68, 68, 90, 48, 48, 49, 34, 44, 34, 73, 115, 115, 117, 101, 65, 116, 34, 58, 48, 44, 34, 67, 97, 116, 101, 103, 111, 114, 121, 34, 58, 34, 34, 125}
	fmt.Println(string(x))
	newgift := &gifts.Gift{}
	err := json.Unmarshal(x, newgift)
	fmt.Println(newgift, err)
}
