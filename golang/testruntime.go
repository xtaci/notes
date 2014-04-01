package main

import (
	"fmt"
	"os"
)

func main() {
	f, e := os.Open("/proc/self/exe")
	if e != nil {
		fmt.Println(f.Stat())
	}
}
