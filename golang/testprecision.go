package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(float32(now.UnixNano()))
	fmt.Println(now.UnixNano())
}
