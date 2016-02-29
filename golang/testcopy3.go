package main

func main() {
	var x []byte
	var y []byte
	println(copy(x, y))
	println(copy([]byte{1}, y))
}
