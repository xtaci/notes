package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	priv, err := rsa.GenerateKey(rand.Reader, 256)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", priv)
	msg := []byte("hello")
	out, err := rsa.EncryptPKCS1v15(rand.Reader, &priv.PublicKey, msg)
	if err != nil {
		panic(err)
	}
	plaintext, err := rsa.DecryptPKCS1v15(nil, priv, out)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(plaintext))
}
