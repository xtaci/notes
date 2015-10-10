package main

import (
	"compress/zlib"
	"io"
	"log"
	"os"
)

func main() {
	r, err := zlib.NewReader(os.Stdin)
	if err != nil {
		log.Println(err)
	}
	io.Copy(os.Stdout, r)
	r.Close()
}
