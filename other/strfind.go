package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
)

type entry struct {
	str string
	ord int // keep words discovery order, since heap is not stable
}

type wordsHeap struct { // only keep those appeared once
	entries []entry
	indices map[string]int  // record the string idx in this heap
	dupstrs map[string]bool // duplicated strings
}

func (h *wordsHeap) Len() int           { return len(h.entries) }
func (h *wordsHeap) Less(i, j int) bool { return h.entries[i].ord < h.entries[j].ord }

func (h *wordsHeap) Swap(i, j int) {
	h.entries[i], h.entries[j] = h.entries[j], h.entries[i]
	h.indices[h.entries[i].str] = i
	h.indices[h.entries[j].str] = j
}

func (h *wordsHeap) Push(x interface{}) {
	h.entries = append(h.entries, x.(entry))
	n := len(h.entries)
	h.indices[x.(entry).str] = n - 1
}

func (h *wordsHeap) Pop() interface{} {
	n := len(h.entries)
	x := h.entries[n-1]
	delete(h.indices, x.str)
	h.entries = h.entries[0 : n-1]
	return x
}

func (h *wordsHeap) Add(str string, ord int) {
	// just ignore the duplicated string
	if h.dupstrs[str] {
		return
	}

	if idx, ok := h.indices[str]; ok {
		// if the string is duplicated, pop that from the heap
		// and record this string in dupstrs
		h.entries[idx].ord = -1
		heap.Fix(h, idx)
		heap.Pop(h)
		h.dupstrs[str] = true
	} else {
		heap.Push(h, entry{str, ord})
	}
}

func (h *wordsHeap) init() {
	h.indices = make(map[string]int)
	h.dupstrs = make(map[string]bool)
}

// findUnique reads from r with a specified bufsize
// and trys to find the first unique string in this file
func findUnique(r io.Reader, bufsize int) {
	reader := bufio.NewReaderSize(r, bufsize)
	h := wordsHeap{}
	h.init()
	ord := 0
	for {
		if line, err := reader.ReadString(' '); err == nil {
			h.Add(string(line), ord)
			ord++
		} else {
			break
		}
	}
	if h.Len() > 0 {
		d := h.Pop().(entry)
		fmt.Println("The first discovered unique word:", d.str)
	} else {
		fmt.Println("Unique words not found!")
	}
	return
}
