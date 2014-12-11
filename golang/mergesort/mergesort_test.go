package mergesort

import (
	//	"sort"
	//	"fmt"
	"helper"
	"runtime"
	"testing"
)

const N = 100000000

func TestMergesort(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	slice := make([]int, N)
	for i := 0; i < N; i++ {
		slice[i] = int(<-helper.LCG)
	}
	println("generated")
	Mergesort(slice, func(a, b int) bool {
		if a < b {
			return true
		}

		return false
	})

	//	fmt.Println(slice)
}

/*
func TestQuicksort(t *testing.T) {
	slice := make([]int, N)
	for i := 0; i < N; i++{
		slice[i] = rand.Int()
	}

	sort.Sort(sort.IntSlice(slice))
//	fmt.Println(slice)
}
*/
