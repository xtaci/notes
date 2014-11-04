package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Add(time.Minute))
	fmt.Println(now)
}
