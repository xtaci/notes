package main

import "testing"

func BenchmarkCopy(b *testing.B) {
	x := make([]byte, 128*1024*1024)
	y := make([]byte, 128*1024*1024)
	b.SetBytes(128 * 1024 * 1024)
	for i := 0; i < b.N; i++ {
		copy(x, y)
	}
}
