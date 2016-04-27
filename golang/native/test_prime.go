package main

import (
	"fmt"
	"math/big"
)

func main() {
	DH1PRIME, _ := big.NewInt(0).SetString("65536", 0)
	fmt.Println(DH1PRIME.BitLen())

	fmt.Println(DH1PRIME.Bytes())
	x := make([]byte, 128)
	copy(x[128-3:], DH1PRIME.Bytes())
	y := new(big.Int).SetBytes(x)
	fmt.Println(y)
}
