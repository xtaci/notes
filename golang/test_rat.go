package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math/big"
)

func main() {
	a := big.NewRat(1, 3)
	b := big.NewRat(7, 22)

	fmt.Println(a, b)
	fmt.Println(a.Float64())

	z := new(big.Rat)
	fmt.Println(z.Add(a, b))
	txt, _ := json.Marshal(a)
	fmt.Println(string(txt))
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf.Bytes())
	dec := gob.NewDecoder(&buf)
	un_a := new(big.Rat)
	dec.Decode(&un_a)
	fmt.Println("un:", un_a)
}
