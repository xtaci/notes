package main

import (
	"encoding/json"
	"log"
)

type Container struct {
	A int32
	F json.RawMessage
}

func main() {
	var c Container
	rawj := `{"a":1, "b":2, "c":"hello world"}`
	c.F = []byte(rawj)

	bts, _ := json.Marshal(c)
	log.Println(string(bts))
	var c2 Container
	err := json.Unmarshal(bts, &c2)
	log.Println(c2, err)
	log.Println(string(c2.F))

	m := make(map[string]json.RawMessage)
	m["rawj"] = []byte(rawj)

	mbts, _ := json.Marshal(m)
	log.Println(string(mbts))
}
