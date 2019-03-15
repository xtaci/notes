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

////////////////////////////////////////////////////////////////////////////////
// pre-processing stage
// a sorter for stream data
type entry struct {
	str []byte
	ord int64
	cnt int64 // string count
}

// split large slice set into a group of small
// sets to mitigate slow runtime.growslice ops
type entrySet struct {
	entries []entry
}

func (s *entrySet) Len() int           { return len(s.entries) }
func (s *entrySet) Less(i, j int) bool { return bytes.Compare(s.entries[i].str, s.entries[j].str) < 0 }
func (s *entrySet) Swap(i, j int)      { s.entries[i], s.entries[j] = s.entries[j], s.entries[i] }

type entrySetReader struct {
	entries []entry
	head    int
}

func (esr *entrySetReader) next() error {
	esr.head++
	if esr.head >= len(esr.entries) {
		return io.EOF
	}
	return nil
}

// memory based aggregator
type memSortAggregator struct {
	sets []entrySetReader
}

func (h *memSortAggregator) Len() int { return len(h.sets) }
func (h *memSortAggregator) Less(i, j int) bool {
	return bytes.Compare(h.sets[i].entries[h.sets[i].head].str, h.sets[j].entries[h.sets[j].head].str) < 0
}
func (h *memSortAggregator) Swap(i, j int)      { h.sets[i], h.sets[j] = h.sets[j], h.sets[i] }
func (h *memSortAggregator) Push(x interface{}) { h.sets = append(h.sets, x.(entrySetReader)) }
func (h *memSortAggregator) Pop() interface{} {
	n := len(h.sets)
	x := h.sets[n-1]
	h.sets = h.sets[0 : n-1]
	return x
}

// words sorter for big memory
type sortWords struct {
	sets     []entrySet
	nextElem int
	setSize  int
	pool     []byte
	offset   int
}

func (h *sortWords) Len() int { return h.nextElem }

func (h *sortWords) Serialize(w io.Writer) {
	if len(h.sets) > 0 {
		agg := new(memSortAggregator)
		for k := range h.sets {
			log.Println("sorting sets#", k)
			sort.Sort(&h.sets[k])
			heap.Push(agg, entrySetReader{h.sets[k].entries, 0})
		}
		log.Println("merging sorted sets to file")

		bufw := bufio.NewWriter(w)
		written := 0
		esr := heap.Pop(agg).(entrySetReader)
		last := esr.entries[esr.head]
		if esr.next() == nil {
			heap.Push(agg, esr)
		}

		for agg.Len() > 0 {
			esr = heap.Pop(agg).(entrySetReader)
			elem := &esr.entries[esr.head]
			if bytes.Compare(elem.str, last.str) == 0 { // condense output
				last.cnt += elem.cnt
			} else {
				bufw.Write(last.str)
				fmt.Fprintf(bufw, ",%v,%v\n", last.ord, last.cnt)
				last = *elem
				written++
			}
			if esr.next() == nil {
				heap.Push(agg, esr)
			}
		}
		bufw.Write(last.str)
		fmt.Fprintf(bufw, ",%v,%v\n", last.ord, last.cnt)
		written++
		log.Println("written", written, "elements")
		bufw.Flush()

		h.offset = 0
		h.nextElem = 0
		h.sets = h.sets[0:0]
	}
}

func (h *sortWords) Add(line []byte, ord int64) bool {
	sz := len(line)
	if h.offset+sz < cap(h.pool) { // limit memory
		copy(h.pool[h.offset:], line)
		if h.nextElem%h.setSize == 0 { // create new set
			entries := make([]entry, 0, h.setSize)
			h.sets = append(h.sets, entrySet{entries})
		}
		idx := h.nextElem / h.setSize
		h.sets[idx].entries = append(h.sets[idx].entries, entry{h.pool[h.offset : h.offset+sz], ord, 1})
		h.offset += sz
		h.nextElem++
		return true
	}
	return false
}

func (h *sortWords) init(limit int64) {
	h.pool = make([]byte, limit)
	h.setSize = 1 << 20 // slice limited to 1M elements
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

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if !h.Add(scanner.Bytes(), ord) {
			fileDump(h, fmt.Sprintf("part%v.dat", parts))
			log.Println("chunk#", parts, "written")
			parts++
			h.Add(scanner.Bytes(), ord)
		}
		ord++

	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error reading from source")
	}

	if h.Len() > 0 {
		fileDump(h, fmt.Sprintf("part%v.dat", parts))
		log.Println("chunk#", parts, "written")
		parts++
	}
	return parts
}

////////////////////////////////////////////////////////////////////////////////
// disk streaming stage
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

// streamAggregator always pop the min string
type streamAggregator struct {
	entries []*streamReader
}

func (h *streamAggregator) Len() int           { return len(h.entries) }
func (h *streamAggregator) Less(i, j int) bool { return h.entries[i].str < h.entries[j].str }
func (h *streamAggregator) Swap(i, j int)      { h.entries[i], h.entries[j] = h.entries[j], h.entries[i] }
func (h *streamAggregator) Push(x interface{}) { h.entries = append(h.entries, x.(*streamReader)) }
func (h *streamAggregator) Pop() interface{} {
	n := len(h.entries)
	x := h.entries[n-1]
	h.entries = h.entries[0 : n-1]
	return x
}

func merger(parts int) chan entry {
	ch := make(chan entry, 4096)
	go func() {
		files := make([]*os.File, parts)
		h := new(streamAggregator)
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
