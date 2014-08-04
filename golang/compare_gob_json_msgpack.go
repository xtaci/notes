package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/vmihailenco/msgpack"
	"time"
)

func main() {

	xxx := make(map[string]interface{})
	for i := int32(0); i < 100000; i++ {
		xxx[fmt.Sprint(i)] = i
	}

	json_buf, _ := json.Marshal(xxx)
	msgpack_buf, _ := msgpack.Marshal(xxx)

	var output bytes.Buffer
	enc := gob.NewEncoder(&output)
	dec := gob.NewDecoder(&output)
	enc.Encode(xxx)
	msg := make(map[string]interface{})

	begin := time.Now()
	dec.Decode(&msg)
	end := time.Now()
	fmt.Println(end.Sub(begin))

	json_doc := make(map[string]interface{})
	begin = time.Now()
	json.Unmarshal(json_buf, &json_doc)
	end = time.Now()
	fmt.Println(end.Sub(begin))

	msgpack_doc := make(map[string]interface{})
	begin = time.Now()
	msgpack.Unmarshal(msgpack_buf, &msgpack_doc)
	end = time.Now()
	fmt.Println(end.Sub(begin))
}
