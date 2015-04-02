package main

func f() (r int) {
	defer func(rr int) {
		r = rr + 5
	}(r)
	return 1
}

func main() {
	println(f())
}
