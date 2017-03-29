package main

import (
	"sync"
	"testing"
)

func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deferfunc()
	}
}

func deferfunc() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
}

func BenchmarkNoDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nodeferfunc()
	}
}

func nodeferfunc() {
	var mu sync.Mutex
	mu.Lock()
	mu.Unlock()
}

func BenchmarkDeferInc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deferInc()
	}
}

func deferInc() {
	i := 0
	defer func() { i++ }()
}

func BenchmarkNoDeferInc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noDeferInc()
	}
}

func noDeferInc() {
	i := 0
	i++
}
