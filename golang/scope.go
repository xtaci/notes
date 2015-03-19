package main

func main() {
	x := 1

	{
		x := 2
		println(x)
	}
	println(x)
}
