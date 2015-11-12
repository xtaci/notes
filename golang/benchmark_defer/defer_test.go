package main

import (
	"sync"
	"testing"
)

var mu sync.Mutex

func deferlock() {
	mu.Lock()
	defer mu.Unlock()
}

func nodeferlock() {
	mu.Lock()
	mu.Unlock()
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
