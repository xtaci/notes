package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.Quote(`hello\t`)
	fmt.Println(s)
	t, err := strconv.Unquote(s)
	fmt.Println(t, err)
}
