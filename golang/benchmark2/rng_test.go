package main

import (
	"crypto/rand"
	"io"
	"testing"
)

func BenchmarkCryptoRead(b *testing.B) {
	buf := make([]byte, 16)
	b.SetBytes(16)
	for i := 0; i < b.N; i++ {
		io.ReadFull(rand.Reader, buf)
	}
}
