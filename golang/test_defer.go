package main

// 函数返回的过程是这样的：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中。

func f() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	println(f())
}
