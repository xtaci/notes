package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"unsafe"
)

////////////////////////////////////////////////////////////////////////////////
// pre-processing stage
// an in-memory sorter
type entry struct {
	str string
	ord int64
	cnt int64 // string count
}

// split large slice set into a group of small
// sets to mitigate slow runtime.growslice ops
type entrySet struct {
	entries []entry
	length  int
}

func (s *entrySet) Len() int           { return s.length }
func (s *entrySet) Less(i, j int) bool { return s.entries[i].str < s.entries[j].str }
func (s *entrySet) Swap(i, j int)      { s.entries[i], s.entries[j] = s.entries[j], s.entries[i] }

type entrySetReader struct {
	entries []entry
	head    int
}

func (esr *entrySetReader) next() bool {
	esr.head++
	if esr.head >= len(esr.entries) {
		return false
	}
	return true
}

// memory based aggregator
type memSortAggregator struct {
	sets []*entrySetReader
}

func (h *memSortAggregator) Len() int { return len(h.sets) }
func (h *memSortAggregator) Less(i, j int) bool {
	return h.sets[i].entries[h.sets[i].head].str < h.sets[j].entries[h.sets[j].head].str
}
func (h *memSortAggregator) Swap(i, j int)      { h.sets[i], h.sets[j] = h.sets[j], h.sets[i] }
func (h *memSortAggregator) Push(x interface{}) { h.sets = append(h.sets, x.(*entrySetReader)) }
func (h *memSortAggregator) Pop() interface{} {
	n := len(h.sets)
	x := h.sets[n-1]
	h.sets = h.sets[0 : n-1]
	return x
}

// words sorter for big memory
type sortWords struct {
	sets        []entrySet
	nextElem    int64
	setMemSize  int64 // memory size of a set
	setSize     int64
	stringUsage int64 // string memory usage track
	setUsage    int64 // set memory usage track
	limit       int64 // max total memory usage
}

func (h *sortWords) Len() int { return int(h.nextElem) }

func (h *sortWords) Serialize(w io.Writer) {
	if len(h.sets) > 0 {
		agg := new(memSortAggregator)
		for k := range h.sets {
			log.Println("sorting sets#", k)
			sort.Sort(&h.sets[k])
			heap.Push(agg, &entrySetReader{h.sets[k].entries[:h.sets[k].length], 0})
		}
		log.Println("merging sorted sets to file")

		written := 0
		esr := heap.Pop(agg).(*entrySetReader)
		last := esr.entries[esr.head]
		if esr.next() {
			heap.Push(agg, esr)
		}

		for agg.Len() > 0 {
			esr = heap.Pop(agg).(*entrySetReader)
			elem := &esr.entries[esr.head]
			if elem.str == last.str { // condense output
				last.cnt += elem.cnt
			} else {
				fmt.Fprintf(w, "%v,%v,%v\n", last.str, last.ord, last.cnt)
				last = *elem
				written++
			}
			if esr.next() {
				heap.Push(agg, esr)
			}
		}
		fmt.Fprintf(w, "%v,%v,%v\n", last.str, last.ord, last.cnt)
		written++
		log.Println("written", written, "elements")

		h.nextElem = 0
		h.sets = nil
		h.setUsage = 0
		h.stringUsage = 0
		runtime.GC()
	}
}

// Add controls the memory for every input
func (h *sortWords) Add(line []byte, ord int64) bool {
	if h.nextElem%h.setSize == 0 { // create new set
		if h.setUsage+h.stringUsage+h.setMemSize > h.limit { // check new set creation
			return false
		}
		entries := make([]entry, h.setSize)
		h.sets = append(h.sets, entrySet{entries, 0})
		h.setUsage += h.setMemSize
	}

	if h.setUsage+h.stringUsage+int64(len(line)) > h.limit { // check new(string)
		return false
	}

	h.stringUsage += int64(len(line))
	sidx := h.nextElem / h.setSize
	eidx := h.nextElem % h.setSize
	h.sets[sidx].entries[eidx] = entry{string(line), ord, 1}
	h.sets[sidx].length++
	h.nextElem++
	return true
}

func (h *sortWords) init(limit int64) {
	e := entry{}
	h.setSize = 1 << 20 // single set limited to 1M elements
	h.setMemSize = int64(unsafe.Sizeof(e)) * h.setSize
	h.limit = limit
	if h.limit < h.setMemSize {
		h.limit = 2 * h.setMemSize
	}
}

// sort2Disk writes strings with it's ordinal and count
// xxxxx,1234,1
// aaaa,5678,10
func sort2Disk(r io.Reader, memLimit int64) int {
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

	scanner := bufio.NewScanner(r)
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
	scanner *bufio.Scanner
	str     string // the head element
	ord     string
	cnt     string
}

func (sr *streamReader) next() bool {
	if sr.scanner.Scan() {
		strs := strings.Split(sr.scanner.Text(), ",")
		sr.str = strs[0]
		sr.ord = strs[1]
		sr.cnt = strs[2]
		return true
	}
	return false
}

func newStreamReader(r io.Reader) *streamReader {
	sr := new(streamReader)
	sr.scanner = bufio.NewScanner(r)
	sr.scanner.Split(bufio.ScanLines)
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
			ch <- entry{sr.str, ord, cnt}
			if sr.next() {
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
		if last_str == e.str {
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
