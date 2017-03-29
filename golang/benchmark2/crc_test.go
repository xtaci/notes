package main

import (
	"hash/crc32"
	"testing"
)

func BenchmarkCRC32(b *testing.B) {
	content := make([]byte, 1024)
	b.SetBytes(int64(len(content)))
	for i := 0; i < b.N; i++ {
		crc32.ChecksumIEEE(content)
	}
}
