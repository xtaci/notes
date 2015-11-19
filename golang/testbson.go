package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type X struct {
	F [4096]byte
}
type S struct {
	A int
	B int
	C int
	D int
}
type X1 struct {
	V [2]S
}

type X2 struct {
	A1 int
	A2 int
	B1 int
	B2 int
	C1 int
	C2 int
	D1 int
	D2 int
}

type X3 struct {
	V [8]int
}

func main() {
	x1 := X1{}
	start := time.Now()
	for i := 0; i < 1000; i++ {
		bson.Marshal(x1)
	}
	fmt.Println(time.Now().Sub(start))

	x2 := X2{}
	start = time.Now()
	for i := 0; i < 1000; i++ {
		bson.Marshal(x2)
	}
	fmt.Println(time.Now().Sub(start))

	x3 := X3{}
	start = time.Now()
	for i := 0; i < 1000; i++ {
		bson.Marshal(x3)
	}
	fmt.Println(time.Now().Sub(start))
}
