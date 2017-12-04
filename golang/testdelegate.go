package main

func A(a int) {
	println("A")
}

func B(b string) {
	println("B")
}

func main() {
	f := A
	f(1)
	f = B
}
