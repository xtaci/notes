package main

func main() {
	var a int
	f(&a)
}

func f(a *int) {
	println(a)
}
