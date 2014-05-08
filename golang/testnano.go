package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(uint32(time.Now().UnixNano()))
}
