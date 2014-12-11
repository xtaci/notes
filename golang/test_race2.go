package main

var c = make(chan int)

//var c = make(chan int, 1)
var a string

func f() {
	a = "hello, world"
	<-c
}

func main() {
	go f()
	c <- 0
	print(a)
}
