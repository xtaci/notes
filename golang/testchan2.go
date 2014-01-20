package main

func main() {
	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
	for v := range ch {
		println(v)
	}
}
