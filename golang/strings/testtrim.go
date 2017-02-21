package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("r:", strings.TrimRight("abcdefedcba", "ab"))
	fmt.Println("l:", strings.TrimLeft("abcdefedcba", "abdef"))
}
