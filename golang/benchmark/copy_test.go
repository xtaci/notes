package main

import (
	"testing"
)

func BenchmarkCopy(b *testing.B) {
	dest := make([]byte, 10, 1024)
	src := make([]byte, 10, 1024)

	for i := 0; i < b.N; i++ {
		copy(dest, src)
	}
}
