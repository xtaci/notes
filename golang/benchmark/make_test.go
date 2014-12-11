package main

import (
	"testing"
	. "types"
)

func BenchmarkMake(b *testing.B) {
	var x *IPCObject
	for i := 0; i < b.N; i++ {
		x = new(IPCObject)
	}
	x.DestID = 1
}
