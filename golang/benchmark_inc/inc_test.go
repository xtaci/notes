package main

import (
	"sync/atomic"
	"testing"
)

var x int32

func BenchmarkInc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomic.AddInt32(&x, 1)
	}
}
