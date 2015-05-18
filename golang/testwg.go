package main

import (
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		go func(idx int) {
			wg.Add(1)
			defer wg.Done()
			print("done")
		}(i)
	}

	wg.Wait()

	print("benchmark completed.\n")
}
