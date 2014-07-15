package main

func add([]float32, []float32) {

}

func main() {
	a := [2]float32{}
	b := [3]float32{}
	add(a[:], b[:])

	x := make([]float32, 3)
	copy(x, a[:])

	str := "abc"
	println(str[:2])
}
