package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"math/big"
)

func main() {
	pri, err := rsa.GenerateKey(rand.Reader, 16)
	if err != nil {
		panic(err)
	}

	fmt.Println("private D:", pri.D)
	fmt.Println("public N:", pri.PublicKey.N, "E:", pri.PublicKey.E)

	const a = 1024
	pubkey1 := big.NewInt(0)
	pubkey2 := big.NewInt(0)
	d1 := big.NewInt(100)
	d2 := big.NewInt(0)
	d2 = d2.Sub(pri.D, d1)
	pubkey1.Exp(big.NewInt(a), d1, pri.PublicKey.N)
	pubkey2.Exp(big.NewInt(a), d2, pri.PublicKey.N)

	fmt.Println(pubkey1)
	fmt.Println(pubkey2)

	ca := big.NewInt(0)
	ca.Exp(big.NewInt(0).Mul(pubkey1, pubkey2), big.NewInt(int64(pri.PublicKey.E)), pri.PublicKey.N)
	fmt.Println(ca)
}
