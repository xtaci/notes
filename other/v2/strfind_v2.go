package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// a heap sorter for stream data
type entry struct {
	str string
	ord int64
}

type wordsHeap struct {
	entries []entry
	memsize int64
}

func (h *wordsHeap) Len() int           { return len(h.entries) }
func (h *wordsHeap) Less(i, j int) bool { return h.entries[i].str < h.entries[j].str }
func (h *wordsHeap) Swap(i, j int)      { h.entries[i], h.entries[j] = h.entries[j], h.entries[i] }
func (h *wordsHeap) Push(x interface{}) { h.entries = append(h.entries, x.(entry)) }
func (h *wordsHeap) Pop() interface{} {
	n := len(h.entries)
	x := h.entries[n-1]
	h.entries = h.entries[0 : n-1]
	return x
}

func (h *wordsHeap) Serialize(w io.Writer) {
	bufw := bufio.NewWriter(w)
	for h.Len() > 0 {
		e := heap.Pop(h).(entry)
		fmt.Fprintf(bufw, "%v,%v\n", e.str, e.ord)
	}
	bufw.Flush()
	h.memsize = 0
}

func (h *wordsHeap) MemSize() int64 { return h.memsize }

func (h *wordsHeap) Add(line string, ord int64) {
	heap.Push(h, entry{line, ord})
	h.memsize = h.memsize + int64(len(line)) + 8 // estimated memory consumption
}

// sort2Disk writes strings with it's ordinal
// xxxxx,1234
// aaaa,5678
func sort2Disk(r io.Reader, memLimit int64) int {
	reader := bufio.NewReader(r)
	h := new(wordsHeap)
	var ord int64
	parts := 0

	// file based serialization
	fileDump := func(hp *wordsHeap, path string) {
		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			log.Fatal(err)
		}
		hp.Serialize(f)
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}

	for {
		line, err := reader.ReadString(' ')
		line = strings.TrimSpace(line)

		if line != "" {
			if h.MemSize()+int64(len(line)) > memLimit {
				fileDump(h, fmt.Sprintf("part%v.dat", parts))
				log.Println("chunk#", parts, "written")
				parts++
			}
			h.Add(string(line), ord)
			ord++
		}

		if err != nil {
			break
		}
	}

	if h.Len() > 0 {
		fileDump(h, fmt.Sprintf("part%v.dat", parts))
		log.Println("chunk#", parts, "written")
		parts++
	}
	return parts
}

type streamReader struct {
	r   *bufio.Reader
	str string // the head element
	ord string
}

func (sr *streamReader) next() error {
	if line, err := sr.r.ReadString('\n'); err == nil {
		line = strings.TrimSpace(line)
		strs := strings.Split(line, ",")
		sr.str = strs[0]
		sr.ord = strs[1]
		return nil
	} else {
		return err
	}
}

func newStreamReader(r io.Reader) *streamReader {
	sr := new(streamReader)
	sr.r = bufio.NewReader(r)
	return sr
}

// pickHeap always pop the min string
type pickHeap struct {
	entries []*streamReader
}

func (h *pickHeap) Len() int           { return len(h.entries) }
func (h *pickHeap) Less(i, j int) bool { return h.entries[i].str < h.entries[j].str }
func (h *pickHeap) Swap(i, j int)      { h.entries[i], h.entries[j] = h.entries[j], h.entries[i] }
func (h *pickHeap) Push(x interface{}) { h.entries = append(h.entries, x.(*streamReader)) }
func (h *pickHeap) Pop() interface{} {
	n := len(h.entries)
	x := h.entries[n-1]
	h.entries = h.entries[0 : n-1]
	return x
}

func merger(parts int) chan entry {
	ch := make(chan entry, 4096)
	go func() {
		files := make([]*os.File, parts)
		h := new(pickHeap)
		for i := 0; i < parts; i++ {
			f, err := os.Open(fmt.Sprintf("part%v.dat", i))
			if err != nil {
				log.Fatal(err)
			}
			files[i] = f
			sr := newStreamReader(f)
			sr.next() // fetch first string,ord
			heap.Push(h, sr)
		}

		for h.Len() > 0 {
			sr := heap.Pop(h).(*streamReader)
			ord, _ := strconv.ParseInt(sr.ord, 10, 64)
			ch <- entry{sr.str, ord}
			if sr.next() == nil {
				heap.Push(h, sr)
			}
		}
		close(ch)

		for _, f := range files[:] {
			if err := f.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}()

	return ch
}

// findUnique reads from r with a specified bufsize
// and trys to find the first unique string in this file
func findUnique(r io.Reader, memLimit int64) {
	// step.1 sort into file chunks
	parts := sort2Disk(r, memLimit)
	log.Println("generated", parts, "parts")
	// step2. sequential output of all parts
	ch := merger(parts)
	log.Println("beginning merged sequential output")

	// step3. loop through the sorted string chan
	// and find the unique string with lowest ord
	var target_str string
	var target_ord int64
	var hasSet bool

	var last_str string
	var last_ord int64
	var last_cnt int
	if e, ok := <-ch; ok {
		last_str = e.str
		last_ord = e.ord
		last_cnt = 1
	} else {
		log.Println("empty set")
		return
	}

	compareTarget := func() {
		if last_cnt == 1 {
			// found new unique string, compare with the ordinal
			if !hasSet {
				target_str = last_str
				target_ord = last_ord
				hasSet = true
			} else if last_ord < target_ord {
				target_str = last_str
				target_ord = last_ord
			}
		}
	}

	// read through the sorted string chan
	for e := range ch {
		if last_str == e.str {
			last_cnt++
		} else {
			compareTarget()
			last_str = e.str
			last_ord = e.ord
			last_cnt = 1
		}
	}

	// make sure the final words is considered
	compareTarget()

	if hasSet {
		log.Println("Found the first unique string:", target_str, "appears at:", target_ord)
	} else {
		log.Println("Unique string not found!")
	}
}
