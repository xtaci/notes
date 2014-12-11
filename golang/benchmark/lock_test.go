package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	_lock sync.RWMutex
)

func lock(λ func()) {
	for i := 0; i < 100; i++ {
		_lock.Lock()
		λ()
		_lock.Unlock()
	}
}

func TestLock(t *testing.T) {
	start := time.Now()
	lock(func() {})
	fmt.Println(time.Now().Sub(start))
}

func BenchmarkLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lock(func() {})
	}
}
