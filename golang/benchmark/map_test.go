package main

import (
	"fmt"
	"testing"
)

func BenchmarkMap(b *testing.B) {
	m := make(map[int32]bool)
	for i := 0; i < b.N; i++ {
		m[int32(i)] = true
	}

	var x bool
	for i := 0; i < b.N; i++ {
		x = m[int32(i)]
	}
	fmt.Println(x)
}
