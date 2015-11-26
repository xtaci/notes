package main

import (
	"fmt"
	"sync"
	"time"
)

const N = 100000

func main() {
	var wg sync.WaitGroup
	var chs []chan *sync.WaitGroup
	for i := 0; i < N; i++ {
		chs = append(chs, make(chan *sync.WaitGroup, 1))
		go f(chs[i])
	}
	time.Sleep(2 * time.Second)
	start := time.Now()
	wg.Add(N)
	for i := 0; i < N; i++ {
		chs[i] <- &wg
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}

func f(ch chan *sync.WaitGroup) {
	for {
		wg := <-ch
		foo()
		wg.Done()
	}
}

func foo() {
	var count int
	for i := 0; i < 10; i++ {
		count += i
	}
}
