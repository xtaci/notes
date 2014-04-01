package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	file, err := os.Open("/Users/xtaci/notes/golang/testdefer.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file.Name())
	fi, _ := file.Stat()
	fmt.Println(path.Ext(file.Name()))
	fmt.Println(strings.TrimRight(fi.Name(), path.Ext(file.Name())))
}
