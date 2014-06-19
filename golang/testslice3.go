package main

func main() {
	xx := []int32{10, 20, 30, 40, 12, 3, 4, 5, 5}
	a := xx[:2]
	b := xx[:3]
	println(cap(a), cap(b))
	println(a[:len(xx)])
}
