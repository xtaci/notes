package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func f() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	// Block until a signal is received.
	<-c
	println("signal")
}

func main() {
	go f()
	go f()
	<-time.After(10 * time.Minute)
}
