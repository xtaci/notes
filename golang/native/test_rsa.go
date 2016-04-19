package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {

	priv, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", priv)
}
