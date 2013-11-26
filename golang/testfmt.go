package main

import (
	"log"
	"fmt"
)

func DEBUG(v ...interface{}) {
    log.Printf("\033[1;33m[DEBUG] %v \033[0m\n", fmt.Sprint(v...))
}

func main() {
	DEBUG("XX","DX")
	DEBUG("1","2")
}
