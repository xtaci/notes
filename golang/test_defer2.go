package main

func main() {
	println(test())
}

func test() (ret int) {
	defer func() {
		ret = 100
	}()
	return 8
}
