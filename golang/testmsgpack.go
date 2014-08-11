package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
)

func main() {
	b := []byte{131, 162, 105, 100, 116, 163, 99, 117, 112, 89, 166, 114, 101, 103, 105, 111, 110, 206, 0, 27, 177, 207}
	y := make(map[string]interface{})
	msgpack.Unmarshal(b, &y)
	fmt.Println(y["id"].(int64))
	fmt.Println(y["cup"].(int64))
	fmt.Println(y["region"].(uint64))

	/*
		fmt.Println(y["A"].(int64))
		fmt.Println(y["B"].(int64))
		fmt.Println(y["C"].(int64))
	*/
}
