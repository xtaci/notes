package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type T interface {
	Add()
}

type Int int32

func (i Int) Add() {
	i++
}

type S struct {
	Int    int
	String string
}

func (s S) Add() {}

type T1 map[string]T

func main() {
	var test = make(T1)
	test["value"] = Int(1)
	test["struct"] = S{23, "ffff"}
	// bson
	var test2 = make(T1)
	b, err := bson.Marshal(test)
	err = bson.Unmarshal(b, &test2)
	fmt.Println(test2)
	fmt.Println(err)
}
