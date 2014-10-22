package main

func get() []int {
	println("get called")
	return []int{1, 2, 3, 4, 5}
}

func main() {
	k := 0
	for k = range get() {
		println("loop ", k)
	}
	println(k)
}
