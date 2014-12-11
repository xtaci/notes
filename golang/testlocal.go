package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
	. "types"
)

type X struct {
	D int
}

func main() {
	data, _ := msgpack.Marshal(X{})
	for i := 0; i < 10; i++ {
		var obj *IPCObject
		msgpack.Unmarshal(data, &obj)
		fmt.Printf("%p\n", obj)
	}
}
