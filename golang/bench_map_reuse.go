package main

import (
	"fmt"
	"time"
)

const n = 100
const size = 100

func main() {

	m := make(map[string]string)
	start := time.Now()

	for l := 0; l < n; l++ {
		for i := 0; i < size; i++ {
			m[fmt.Sprint(i)] = fmt.Sprint(i)
		}

		for i := 0; i < size; i++ {
			delete(m, fmt.Sprint(i))
		}
	}

	fmt.Println("delete and reuse:", time.Now().Sub(start))

	start = time.Now()

	for l := 0; l < n; l++ {
		m = make(map[string]string)
		for i := 0; i < size; i++ {
			m[fmt.Sprint(i)] = fmt.Sprint(i)
		}
	}

	fmt.Println("make new one", time.Now().Sub(start))

}
