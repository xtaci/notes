package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 10000000

var (
	_alloc_lock sync.Mutex
	_buf        []byte
	_pos        int
)

func init() {
	_buf = make([]byte, 1024*1024)
}

func main() {
	test_make()
	test_allocbyte()
}

func test_make() {
	start := time.Now()
	for i := 0; i < N; i++ {
		mk()
	}
	fmt.Println(time.Now().Sub(start))
}

func mk() []byte {
	return make([]byte, 1+rand.Int()%100)
}

func test_allocbyte() {
	start := time.Now()
	for i := 0; i < N; i++ {
		allocbyte(1 + rand.Int()%100)
	}
	fmt.Println(time.Now().Sub(start))
}

func allocbyte(size int) []byte {
	//	_alloc_lock.Lock()
	//	defer _alloc_lock.Unlock()
	if _pos+size > len(_buf) {
		_buf = make([]byte, 1024*1024)
		_pos = 0
	}
	ret := _buf[_pos : _pos+size]
	_pos += size
	return ret
}
