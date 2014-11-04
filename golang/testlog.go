package main

import (
	"fmt"
	"strings"
)

func main() {
	ERR("HELLO", "WOLRD")
}

func ERR(v ...interface{}) {
	fmt.Printf("\033[1;7;31m[ERROR] %v \033[0m\n", strings.TrimSpace(fmt.Sprintln(v...)))
}
