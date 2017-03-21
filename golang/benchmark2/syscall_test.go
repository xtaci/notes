package main

import (
	"testing"
	"time"
)

func BenchmarkNow(b *testing.B) {
	var now uint32
	for i := 0; i < b.N; i++ {
		now = currentMs()
	}
	_ = now
}

func currentMs() uint32 {
	return uint32(time.Now().UnixNano() / int64(time.Millisecond))
}
