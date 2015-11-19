package main

func main() {
	x := uint32(0xDEADBEAF)
	println(byte(x))
	println(byte(x >> 8))
	println(byte(x >> 16))
	println(byte(x >> 24))
}
