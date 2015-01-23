package main

import "github.com/davecgh/go-spew/spew"

func main() {
	slice := []byte("line 1\r\n")
	spew.Dump(slice)
}
