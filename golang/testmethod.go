package main

type T int

func (v T) test() {
	println("test")
}

func main() {
	var t T
	(&t).test()
}
