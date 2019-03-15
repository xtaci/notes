package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// a heap sorter for stream data
type entry struct {
	str []byte
	ord int64
	cnt int64 // string count
}

type sortWords struct {
	entries []entry
	pool    []byte
	offset  int
}

func (h *sortWords) Len() int           { return len(h.entries) }
func (h *sortWords) Less(i, j int) bool { return bytes.Compare(h.entries[i].str, h.entries[j].str) < 0 }
func (h *sortWords) Swap(i, j int)      { h.entries[i], h.entries[j] = h.entries[j], h.entries[i] }

func (h *sortWords) Serialize(w io.Writer) {
	if h.Len() > 0 {
		sort.Sort(h)
		bufw := bufio.NewWriter(w)
		last := h.entries[0]
		for k := range h.entries {
			if bytes.Compare(h.entries[k].str, last.str) == 0 { // condense output
				last.cnt += h.entries[k].cnt
			} else {
				bufw.Write(last.str)
				fmt.Fprintf(bufw, ",%v,%v\n", last.ord, last.cnt)
				last = h.entries[k]
			}
		}
		bufw.Write(last.str)
		fmt.Fprintf(bufw, ",%v,%v\n", last.ord, last.cnt)
		bufw.Flush()
		h.offset = 0
		h.entries = h.entries[0:0]
	}
}

func (h *sortWords) Add(line string, ord int64) bool {
	sz := len(line)
	if h.offset+sz < cap(h.pool) { // limit memory
		copy(h.pool[h.offset:], []byte(line))
		h.entries = append(h.entries, entry{h.pool[h.offset : h.offset+sz], ord, 1})
		h.offset += sz
		return true
	}
	return false
}

func (h *sortWords) init(limit int64) {
	h.pool = make([]byte, limit)
}

// sort2Disk writes strings with it's ordinal and count
// xxxxx,1234,1
// aaaa,5678,10
func sort2Disk(r io.Reader, memLimit int64) int {
	reader := bufio.NewReader(r)
	h := new(sortWords)
	h.init(memLimit)
	var ord int64
	parts := 0

	// file based serialization
	fileDump := func(hp *sortWords, path string) {
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
			if !h.Add(line, ord) {
				fileDump(h, fmt.Sprintf("part%v.dat", parts))
				log.Println("chunk#", parts, "written")
				parts++
				h.Add(line, ord)
			}
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
	cnt string
}

func (sr *streamReader) next() error {
	if line, err := sr.r.ReadString('\n'); err == nil {
		line = strings.TrimSpace(line)
		strs := strings.Split(line, ",")
		sr.str = strs[0]
		sr.ord = strs[1]
		sr.cnt = strs[2]
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
			cnt, _ := strconv.ParseInt(sr.cnt, 10, 64)
			ch <- entry{[]byte(sr.str), ord, cnt}
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
	var target_str []byte
	var target_ord int64
	var hasSet bool

	var last_str []byte
	var last_ord int64
	var last_cnt int64
	if e, ok := <-ch; ok {
		last_str = e.str
		last_ord = e.ord
		last_cnt = e.cnt
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
		if bytes.Compare(last_str, e.str) == 0 {
			last_cnt += e.cnt
		} else {
			compareTarget()
			last_str = e.str
			last_ord = e.ord
			last_cnt = e.cnt
		}
	}

	// make sure the final words is considered
	compareTarget()

	if hasSet {
		log.Println("Found the first unique string:", string(target_str), "appears at:", target_ord)
	} else {
		log.Println("Unique string not found!")
	}
}
