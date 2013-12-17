package main

import (
	"fmt"
	"log"
	"strings"
)

func DEBUG(v ...interface{}) {
	log.Printf("\033[1;33m[DEBUG] %v \033[0m\n", strings.TrimRight(fmt.Sprintln(v...), "\n"))
}

func main() {
	DEBUG("XX", "DX")
	DEBUG("1", "2")
	DEBUG("weapon", "level", 18, "armor_level", 18)
	fmt.Println("weapon", "level", 18, "armor_level", 18)
}
