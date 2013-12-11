package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {
	xxx := make(map[int32]int32)
	for i:=int32(0);i<10;i++ {
		xxx[i] = i
	}

	var output bytes.Buffer
	enc := gob.NewEncoder(&output)
	dec := gob.NewDecoder(&output)
	enc.Encode(xxx)
	fmt.Println(output)
	msg := make(map[int32]int32)
	dec.Decode(&msg)
	fmt.Println(msg)
}
