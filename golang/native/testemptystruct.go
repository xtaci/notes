package main

import "time"

func main() {
	ch := make(chan struct{}, 1024*1024*1024*1024)
	for i := 0; i < 1024*1024*1024*1024; i++ {
		ch <- struct{}{}
	}
	time.Sleep(time.Minute)
}
