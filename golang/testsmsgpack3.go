package main

import (
	"fmt"
	"gopkg.in/vmihailenco/msgpack.v2"
)

type Inner struct {
	A int
	B int
}

type Word struct {
	Words string
	In    interface{}
}

func main() {
	w := Word{Words: "HELLO", In: Inner{10, 20}}
	b, err := msgpack.Marshal(w)
	fmt.Println(string(b), err)

	n := Word{}
	err = msgpack.Unmarshal(b, &n)
	fmt.Println(n, err)
}
