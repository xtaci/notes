package main

import (
	"time"
)

func fun(f func()) {
	for i := 0; i < 10; i++ {
		f()
	}
}

func fun2() {
	x := 1
	go fun(func() {
		x++
		println(x)
	})
}

func main() {
	fun2()
	time.Sleep(10)
}
