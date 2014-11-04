package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	const N = 40000

	arr := make([]string, 0, N)
	for i := 0; i < N; i++ {
		arr = append(arr, fmt.Sprint(i))
	}

	start := time.Now()
	for i := 0; i < N; i++ {
		strconv.Atoi(arr[i])
	}

	fmt.Println(time.Now().Sub(start))
}
