package main

var c = make(chan int, 1)
var a string

func f() {
	a = "hello, world"
	for {
		<-c
	}
}
func main() {
	go f()
	for {
		c <- 0
		print(a)
	}
}
