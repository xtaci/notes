package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"reflect"
)

var str = `
{
    "a": 134,
    "b": 1.111111,
    "c": 4294967296,
    "d": "abc",
    "e": 4294967296.999999,
    "f": 65536
}
`

func main() {
	var x map[string]interface{}
	var y map[string]interface{}
	x = make(map[string]interface{})
	x["a"] = int(134)
	x["b"] = float32(1.111111)
	x["c"] = int64(4294967296)
	x["d"] = "abc"
	x["e"] = float64(4294967296.999999)
	x["f"] = uint16(65535)
	x["g"] = int32(65535)
	x["h"] = 98765432100000

	data, err := bson.Marshal(&x)
	if err != nil {
		fmt.Println(err)
	}

	bson.Unmarshal(data, &y)
	for k, v := range y {
		fmt.Println(k, v, reflect.ValueOf(v))
	}

	var z interface{}
	z = 1234
	m, ok := z.(int)
	fmt.Println(m, ok)
}
