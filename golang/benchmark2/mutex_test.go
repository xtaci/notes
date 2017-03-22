package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	var mu sync.Mutex
	var sum int
	for i := 0; i < b.N; i++ {
		mu.Lock()
		sum += i
		mu.Unlock()
	}
}

func BenchmarkAtomic(b *testing.B) {
	var sum int64
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&sum, int64(i))
	}
}
