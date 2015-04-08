package main

var test int

func f(i, j int) func() (int, int) {
	return func() (int, int) {
		i++
		j++
		test++
		return i, j
	}
}
