package main

import (
	"fmt"
	"gopkg.in/vmihailenco/msgpack.v2"
)

type Inner2 struct {
	A int
	B int
}

type Inner struct {
	A int
	B int
	C interface{}
}

type Word struct {
	Words string
	In    Inner
}

func main() {
	w := Word{"HELLO", Inner{10, 20, Inner2{30, 40}}}
	b, err := msgpack.Marshal(w)
	fmt.Println(string(b), err)

	n := make(map[string]interface{})
	err = msgpack.Unmarshal(b, &n)
	fmt.Println(n, err)
}
