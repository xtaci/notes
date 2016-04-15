package main

import (
	"fmt"

	. "github.com/aerospike/aerospike-client-go"
)

type Y struct {
	a int
}

type X struct {
	b int
	y Y
}

func main() {
	// define a client to connect to
	client, err := NewClient("127.0.0.1", 3000)
	panicOnError(err)

	namespace := "test"
	setName := "aerospike"
	key, err := NewKey(namespace, setName, "key") // user key can be of any supported type
	panicOnError(err)

	// define some bins
	x := X{}
	bins := BinMap{
		"bin1": []interface{}{x},
	}

	// write the bins
	writePolicy := NewWritePolicy(0, 0)
	err = client.Put(writePolicy, key, bins)
	panicOnError(err)

	// read it back!
	readPolicy := NewPolicy()
	rec, err := client.Get(readPolicy, key)
	panicOnError(err)

	fmt.Printf("%#v\n", *rec)
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
