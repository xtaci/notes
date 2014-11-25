package main

import (
	"github.com/vmihailenco/msgpack"
	"github.com/davecgh/go-spew/spew"
	"testing"
//	. "types"
)

func BenchmarkMsgPack(b *testing.B) {
	obj := int32(1) // &Words{}
	data, _ := msgpack.Marshal(obj)

	spew.Dump(data)
	var restore int32
	for i := 0; i < b.N; i++ {
		msgpack.Unmarshal(data, &restore)
	}
}
