package main

import (
	"sync"
	"testing"
)

var mu sync.RWMutex

func deferlock() {
	mu.RLock()
	defer mu.RUnlock()
}

func nodeferlock() {
	mu.RLock()
	mu.RUnlock()
}

func BenchmarkDeferLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deferlock()
	}
}

func BenchmarkNoDeferLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nodeferlock()
	}
}
