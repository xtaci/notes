package main

import (
	"misc/packet"
	"testing"
)

type BOOL struct {
	F_v bool
}

func BenchmarkPack(b *testing.B) {
	x := BOOL{}
	for i := 0; i < b.N; i++ {
		packet.Pack(-1, x, nil)
	}
}
