package main

import (
	"log"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		f()
		<-time.After(10 * time.Millisecond)
	}

	select {}
}

func f() {
	ticker := time.NewTicker(10 * time.Millisecond)
	log.Println(ticker)
}
