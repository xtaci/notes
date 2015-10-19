package main

func main() {
	var x map[string]map[string]int
	println(x["a"]["b"])
	delete(x, "a")
	println(len(x))
}
