package main

import (
	//	"os"
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("exit")
	}()

	//	os.Exit(0)
}
