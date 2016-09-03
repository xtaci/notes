package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	l := new(sync.Mutex)
	c := sync.NewCond(l)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("waiting %d\n", i)
			l.Lock()
			defer l.Unlock()
			c.Wait()
			fmt.Printf("go %d\n", i)
		}(i)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		c.Signal()
	}
	time.Sleep(1 * time.Second)
}
