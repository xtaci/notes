package mergesort

import (
	"container/heap"
	"runtime"
	"sync"
	"sort"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	heap.Remove(pq, item.index)
	item.value = value
	item.priority = priority
	heap.Push(pq, item)
}

func _merge(slice []int, left, mid, right int) []int {
	temp := make([]int, right-left+1)
	pos := 0
	lpos := left
	rpos := mid + 1

	for lpos <= mid && rpos <= right {
		if slice[lpos] < slice[rpos] {
			temp[pos] = slice[lpos]
			pos++
			lpos++
		} else {
			temp[pos] = slice[rpos]
			pos++
			rpos++
		}
	}

	for lpos <= mid {
		temp[pos] = slice[lpos]
		pos++
		lpos++
	}

	for rpos <= right {
		temp[pos] = slice[rpos]
		pos++
		rpos++
	}

	return temp
}

func _qs(sli []int, wg *sync.WaitGroup) {
	sort.Sort(sort.IntSlice(sli))
	wg.Done()
}

//------------------------------------------------ 并行的归并排序
// 注意，直接使用的时候最好设置一下GOMAXPROCS，避免调度的问题
func Mergesort(slice []int, less func(a, b int) bool) {
	n := runtime.NumCPU()
	sz := len(slice)/n

	// quicksort first
	wg := &sync.WaitGroup{}
	wg.Add(n)
	i:=0
	for ;i<n-1;i++ {
		go _qs(slice[i*sz : (i+1) * sz], wg)
	}
	go _qs(slice[i*sz:], wg)
	wg.Wait()

	// merge
}
