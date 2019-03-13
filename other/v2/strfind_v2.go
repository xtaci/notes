package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const MaxHeapSize = 1024 * 1024 * 1024

// a heap sorter for stream data
type entry struct {
	str string
	ord int
}

type wordsHeap struct {
	entries []entry
	limit   int
}

func (h *wordsHeap) Len() int { return len(h.entries) }
func (h *wordsHeap) Less(i, j int) bool {
	return strings.Compare(h.entries[i].str, h.entries[j].str) == -1
}

func (h *wordsHeap) Swap(i, j int) { h.entries[i], h.entries[j] = h.entries[j], h.entries[i] }

func (h *wordsHeap) Push(x interface{}) { h.entries = append(h.entries, x.(entry)) }

func (h *wordsHeap) Pop() interface{} {
	n := len(h.entries)
	x := h.entries[n-1]
	h.entries = h.entries[0 : n-1]
	return x
}

func (h *wordsHeap) Add(str string, ord int) bool {
	if len(h.entries) < h.limit {
		heap.Push(h, entry{str, ord})
		return true
	} else {
		return false
	}
}

func (h *wordsHeap) Serialize(w io.Writer) {
	bufw := bufio.NewWriter(w)
	for h.Len() > 0 {
		e := heap.Pop(h).(entry)
		fmt.Fprintf(bufw, "%v,%v\n", e.str, e.ord)
	}
	bufw.Flush()
}

func (h *wordsHeap) init(limit int) { h.limit = limit }

// sort2Disk writes strings with it's ordinal
// xxxxx,1234
// aaaa,5678
func sort2Disk(r io.Reader, bufsize int) int {
	reader := bufio.NewReaderSize(r, bufsize)
	h := wordsHeap{}
	h.init(MaxHeapSize)
	ord := 0
	parts := 0
	for {
		if line, err := reader.ReadString(' '); err == nil {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			if !h.Add(string(line), ord) {
				f, err := os.OpenFile(fmt.Sprintf("part%v.dat", parts), os.O_RDWR|os.O_CREATE, 0755)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(h.Len())
				h.Serialize(f)
				if err := f.Close(); err != nil {
					log.Fatal(err)
				}
				parts++
				h = wordsHeap{}
				h.init(MaxHeapSize)
				h.Add(string(line), ord)
			}
			ord++
		} else {
			break
		}
	}

	f, err := os.OpenFile(fmt.Sprintf("part%v.dat", parts), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	h.Serialize(f)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	return parts
}

// findUnique reads from r with a specified bufsize
// and trys to find the first unique string in this file
func findUnique(r io.Reader, bufsize int) {
	parts := sort2Disk(r, bufsize)
	log.Println(parts)
}
