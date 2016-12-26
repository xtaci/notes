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
	c.F = json.RawMessage(`{"a":1, "b":2, "c":"hello world"}`)

	bts, _ := json.Marshal(c)
	log.Println(string(bts))
	var c2 Container
	err := json.Unmarshal(bts, &c2)
	log.Println(c2, err)
	log.Println(string(c2.F))
}
