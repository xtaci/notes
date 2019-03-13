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
	limit   int64
	memsize int64
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

func (h *wordsHeap) Serialize(w io.Writer) {
	bufw := bufio.NewWriter(w)
	for h.Len() > 0 {
		e := heap.Pop(h).(entry)
		fmt.Fprintf(bufw, "%v,%v\n", e.str, e.ord)
	}
	bufw.Flush()
}
func (h *wordsHeap) MemSize() int64 { return h.memsize }

func (h *wordsHeap) Add(line string, ord int64) {
	heap.Push(h, entry{line, ord})
	h.memsize = h.memsize + int64(len(line)) + 8 // estimated memory consumption
}

func (h *wordsHeap) init(limit int64) { h.limit = limit }
func (h *wordsHeap) Limit() int64     { return h.limit }

// sort2Disk writes strings with it's ordinal
// xxxxx,1234
// aaaa,5678
func sort2Disk(r io.Reader, memLimit int64) int {
	reader := bufio.NewReader(r)
	h := new(wordsHeap)
	h.init(memLimit)
	var ord int64
	parts := 0
	for {
		line, err := reader.ReadString(' ')
		line = strings.TrimSpace(line)

		if line != "" {
			h.Add(string(line), ord)
			ord++

			if h.MemSize() >= h.Limit() {
				f, err := os.OpenFile(fmt.Sprintf("part%v.dat", parts), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
				if err != nil {
					log.Fatal(err)
				}
				h.Serialize(f)
				if err := f.Close(); err != nil {
					log.Fatal(err)
				}
				log.Println("chunk#", parts, "written")
				parts++
				h = new(wordsHeap)
				h.init(memLimit)
			}
		}

		if err != nil {
			break
		}
	}

	f, err := os.OpenFile(fmt.Sprintf("part%v.dat", parts), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	h.Serialize(f)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	log.Println("chunk#", parts, "written")
	parts++
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

func (h *pickHeap) Len() int { return len(h.entries) }
func (h *pickHeap) Less(i, j int) bool {
	return strings.Compare(h.entries[i].str, h.entries[j].str) == -1
}

func (h *pickHeap) Swap(i, j int) { h.entries[i], h.entries[j] = h.entries[j], h.entries[i] }

func (h *pickHeap) Push(x interface{}) { h.entries = append(h.entries, x.(*streamReader)) }

func (h *pickHeap) Pop() interface{} {
	n := len(h.entries)
	x := h.entries[n-1]
	h.entries = h.entries[0 : n-1]
	return x
}

// merger combines all small parts into large sorted file
func merger(parts int, w io.Writer) {
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

	bufw := bufio.NewWriter(w)
	for h.Len() > 0 {
		sr := heap.Pop(h).(*streamReader)
		fmt.Fprintf(bufw, "%v,%v\n", sr.str, sr.ord)
		if sr.next() == nil {
			heap.Push(h, sr)
		}
	}
	bufw.Flush()

	for _, f := range files[:] {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

// findUnique reads from r with a specified bufsize
// and trys to find the first unique string in this file
func findUnique(r io.Reader, memLimit int64) {
	// step.1 sort into file chunks
	parts := sort2Disk(r, memLimit)
	log.Println("generated", parts, "parts")
	f, err := os.OpenFile("big.dat", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// step2. merge all these parts into a sorted large file
	merger(parts, f)
	// step3. loop through the file and find the unique string with lowest ord
	if _, err := f.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	var target_str string
	var target_ord int64
	var hasSet bool

	sr := newStreamReader(f)
	sr.next()
	last_str := sr.str
	last_ord := sr.ord
	last_cnt := 1

	compareTarget := func() {
		if last_cnt == 1 {
			// found new unique string, compare with the ordinal
			if !hasSet {
				target_str = last_str
				target_ord, _ = strconv.ParseInt(last_ord, 10, 64)
				hasSet = true
			} else {
				new_ord, _ := strconv.ParseInt(last_ord, 10, 64)
				if new_ord < target_ord {
					target_str = last_str
					target_ord = new_ord
				}
			}
		}
	}

	for sr.next() == nil {
		if last_str == sr.str {
			last_cnt++
		} else {
			compareTarget()

			// record current
			last_str = sr.str
			last_ord = sr.ord
			last_cnt = 1
		}
	}

	// make sure the final words is considered
	compareTarget()

	if hasSet {
		fmt.Println("Found the first unique string:", target_str, "appears at:", target_ord)
	} else {
		fmt.Println("Unique string not found!")
	}

	// cleanup
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
