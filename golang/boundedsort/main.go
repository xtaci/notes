package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	const N = 1000
	input := make([]int, N)
	for i := 0; i < N; i++ {
		input[i] = i
	}
	fmt.Println("input:", input)

	fmt.Println("output:")
	sort(input)
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h IntHeap) Update(o, n int) {
	for i := range h {
		if h[i] == o {
			h[i] = n
			heap.Fix(&h, i)
			return
		}
	}
}

const bound = 20

func sort(input []int) {
	loopcount := len(input)/bound + 1
	max := math.MaxUint32
	i := 0
	for ; i < loopcount; i++ {
		max = phase(input, max)
	}
}

func phase(input []int, max int) int {
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < len(input); i++ {
		if input[i] < max {
			heap.Push(h, input[i])
		}
		if h.Len() > bound {
			heap.Pop(h)
		}
	}

	if h.Len() > 0 {
		max = heap.Pop(h).(int)
		fmt.Printf("%v ", max)
		for h.Len() > 0 {
			v := heap.Pop(h).(int)
			fmt.Printf("%v ", v)
		}
		fmt.Println()
	}
	return max
}
