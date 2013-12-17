package main

import (
	"fmt"
	"strings"
)

func main() {
	line := "TAX000,0,0,"
	fields := strings.Split(line, ",")
	fmt.Println(strings.Join(fields, "#"), len(fields))
}
