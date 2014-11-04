package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	monthday := time.Date(now.Year(), 3, 0, 0, 0, 0, 0, time.UTC)
	fmt.Println(monthday, monthday.Day())
}
