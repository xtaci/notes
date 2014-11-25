package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
)

type WordOld struct {
	Words  string
}

type WordNew struct {
	Words  string `msgpack:"W"`
}

func main() {
	old := WordNew{"HELLO"}
	b, err := msgpack.Marshal(old)
	fmt.Println(string(b), err)

	n := WordNew{}
	err = msgpack.Unmarshal(b, &n)
	fmt.Println(n,err)
}
