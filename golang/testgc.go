package main

import (
	. "github.com/gonet2/rank/src/rank/dos"
)

func main() {
	tree := Tree{}
	const N = 10000000
	for i := 0; i < N; i++ {
		tree.Insert(int32(i), int32(i))
	}

	/*
		for i := 0; i < N; i++ {
		}*/
	println("load complete")
	i := 0
	for {
		if i == N {
			i = 0
		}
		_, n := tree.Locate(int32(i), int32(i))
		tree.Delete(int32(i), n)
		tree.Insert(int32(i), int32(i))
		i++
	}
}
