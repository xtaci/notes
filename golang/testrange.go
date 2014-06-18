package main

func main() {
	x := []int{1, 2, 3}
	k := 0
	for k = range x {
		println("loop ", k)
	}
	println(k)
}
