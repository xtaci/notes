package main

func main() {
	base := make([]int, 0, 32)
	// index range
	println(base[0])
	// capacity limit
	extend := base[:33]
	println(extend)
}
