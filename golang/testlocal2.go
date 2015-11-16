package main

func main() {
	for i := uint64(0); i < 10; i++ {
		var x uint64
		x |= i
		println(x)
	}
}
