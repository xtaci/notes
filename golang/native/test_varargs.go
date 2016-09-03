package main

func f(a, b int, opts ...int) {
	println(a, b)
	for k := range opts {
		println(opts[k])
	}
}

func main() {
	f(1, 2)
	f(1, 2, 1, 2, 3, 4, 5)
}
