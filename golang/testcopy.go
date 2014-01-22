package main

import (
	"fmt"
	"sync"
	"time"
)

type Words struct {
	Words     string
	SpeakerId int32
	Speaker   string
	Timestamp int64
	VIPLevel  int32
	Lang      string
	TYPE      int8
}

func main() {
	var lock sync.Mutex
	src := make([]Words, 100)
	lock.Lock()
	defer lock.Unlock()
	dest := make([]Words, 1000)

	time.Sleep(time.Second)
	start := time.Now()
	copy(dest, src)
	fmt.Println(len(dest))
	fmt.Println(time.Now().Sub(start))
}
