package main

import (
	"runtime"
	"time"
)

func main() {
	tick()
	runtime.GC()
	select {}
}

func finalizer(t *time.Ticker) {
	println("ticker gc")
}

func finalizer2(t *time.Time) {
	println("time gc")
}

func tick() {
	ticker := time.NewTicker(10 * time.Millisecond)
	t := time.Now()
	runtime.SetFinalizer(ticker, finalizer)
	runtime.SetFinalizer(&t, finalizer2)
}
