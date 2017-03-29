package main

import "testing"

func BenchmarkSelectNotify(b *testing.B) {
	ch := make(chan struct{}, 1)
	for i := 0; i < b.N; i++ {
		select {
		case ch <- struct{}{}:
		default:
		}
	}
}
