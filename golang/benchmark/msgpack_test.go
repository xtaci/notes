package main

import (
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/vmihailenco/msgpack.v2"
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
