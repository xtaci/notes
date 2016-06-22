package main

import (
	"fmt"

	"github.com/klauspost/reedsolomon"
)

func main() {
	enc, err := reedsolomon.New(10, 3)
	data := make([][]byte, 13)
	for i := 0; i < 13; i++ {
		data[i] = make([]byte, 1024)
	}
	for i, in := range data[:10] {
		for j := range in {
			in[j] = byte((i + j) & 0xff)
		}
	}
	fmt.Println(data[0])
	err = enc.Encode(data)
	println(err)
	ok, err := enc.Verify(data)
	println(ok)
	fmt.Println(data[0])
	data[0] = nil
	err = enc.Reconstruct(data)
	println(err)
	fmt.Println("reconstruct:", data[0])
}
