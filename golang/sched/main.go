package main

import "time"

func main() {
	for i := 0; i < 3000; i++ {
		go f()
	}
	select {}
}

func f() {
	tc := time.After(33 * time.Millisecond)

	for {
		<-tc
		tc = time.After(33 * time.Millisecond)
	}
}
