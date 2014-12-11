package main

import (
	"fmt"
	"time"
	. "types"
)

func main() {
	ch := make(chan *IPCObject, 1000)
	go func() {
		for {
			ch <- new(IPCObject)
		}
	}()

	time.Sleep(2 * time.Second)
	start := time.Now()
	//<-ch
	//	x := new(IPCObject)
	fmt.Println(time.Now().Sub(start))
	//	fmt.Println(x)
}
