package main

import (
	"crypto/rand"
	"io"

	"github.com/golang/snappy"
)

func main() {
	data := make([]byte, 1350)
	io.ReadFull(rand.Reader, data[:1300])
	println(len(snappy.Encode(nil, data)))
}
