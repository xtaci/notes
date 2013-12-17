package main

func fun(f func()) {
	for i := 0; i < 10; i++ {
		f()
	}
}

func main() {
	x := 1
	lambda := func() {
		x++
		println(x)
	}
	fun(lambda)
}
