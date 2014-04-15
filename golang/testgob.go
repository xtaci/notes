package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

func main() {

	begin := time.Now()
	xxx := make(map[int32]int32)
	for i := int32(0); i < 1000; i++ {
		xxx[i] = i
	}

	var output bytes.Buffer
	enc := gob.NewEncoder(&output)
	dec := gob.NewDecoder(&output)
	enc.Encode(xxx)
	msg := make(map[int32]int32)
	dec.Decode(&msg)
	end := time.Now()
	fmt.Println(end.Sub(begin))
}
