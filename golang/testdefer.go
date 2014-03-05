package main

func test() (value float64) {
	defer func() {
		value = value * 1.25
	}()

	return 100
}

func main() {
	println(test())
}
