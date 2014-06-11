package main

func test() (value float64) {
	defer func() {
		value = value * 1.25
	}()

	return 100
}

func main() {
	defer func() {
		dealpanic()
	}()

	panic("test")
}

func dealpanic() {
	if x := recover(); x != nil {
		println(x)
	}
}
