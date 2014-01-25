package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := make([]byte, 1)
	bf := bytes.NewBuffer(b)
	fmt.Println(bf.Bytes())
	bf.Write([]byte{1, 2, 3, 4, 5})
	bf.Write([]byte{1, 2, 3, 4, 5})
	fmt.Println(bf.Bytes())
}
