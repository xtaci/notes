package main

import (
	"container/list"
	"testing"
)

func BenchmarkSlice(b *testing.B) {
	s := make([]byte, b.N)
	z := byte(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		z += s[i]
	}
}

func BenchmarkList(b *testing.B) {
	l := list.New()
	for i := 0; i < b.N; i++ {
		l.PushBack(byte(i))
	}

	z := byte(0)
	b.ResetTimer()
	for e := l.Front(); e != nil; e = e.Next() {
		z += (e.Value).(byte)
	}
}
