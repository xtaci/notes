package main

import (
	"fmt"
	"gopkg.in/vmihailenco/msgpack.v2"
)

type WordOld struct {
	w string
}

type WordNew struct {
	w string
}

func main() {
	old := WordNew{"HELLO"}
	b, err := msgpack.Marshal(old)
	fmt.Println(string(b), err)

	n := WordNew{}
	err = msgpack.Unmarshal(b, &n)
	fmt.Println(n, err)
}
