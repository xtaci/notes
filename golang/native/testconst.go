package main

func main() {
	f(10)
}

var a int = 10

const b = int(10)

func f(a int) {
	const b int = int(a)
	println(b)
}
