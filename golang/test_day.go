package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix()
	fmt.Println(time.Unix(now-now%86400, 0))
}
