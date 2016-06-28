package main

func main() {
	for i := 1; i < 100; i++ {
		x := make([]byte, i, i*2)
		println(len(x), cap(x), cap(x[1:]))
	}
}
