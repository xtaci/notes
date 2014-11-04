package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type L struct {
	V    int
	Next *L
}

func main() {

	A := &L{}
	A.Next = A
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(A)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(buf.Bytes())
}
