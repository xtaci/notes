package main

import (
	"fmt"
	"github.com/google/btree"
	"time"
)

func main() {
	for i := 1; i <= 1000000; i *= 10 {
		bench_btree(i)
		bench_map(i)
	}
}

func bench_btree(n int) {
	tr := btree.New(100)
	start := time.Now()
	for i := 0; i < n; i++ {
		tr.ReplaceOrInsert(btree.Int(i))
	}
	fmt.Printf("btree_insert: count:%v time:%v\n", n, time.Now().Sub(start))

	start = time.Now()
	for i := 0; i < n; i++ {
		tr.Get(btree.Int(i))
	}
	fmt.Printf("btree_get: count:%v time:%v\n", n, time.Now().Sub(start))
}

func bench_map(n int) {
	m := make([]int, n)
	start := time.Now()
	for i := 0; i < n; i++ {
		m[i] = i
	}
	fmt.Printf("map_insert: count:%v time:%v\n", n, time.Now().Sub(start))

	start = time.Now()
	for i := 0; i < n; i++ {
		_ = m[i]
	}
	fmt.Printf("map_get: count:%v time:%v\n", n, time.Now().Sub(start))
}
