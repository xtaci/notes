package main

func main() {
	x := new([]int)
	println(cap(*x))
	y := make([]int, 0, 0)
	println(cap(y))
}
