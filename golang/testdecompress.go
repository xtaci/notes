package main

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	raw64 := ""
	data, err := base64.StdEncoding.DecodeString(raw64)
	r := flate.NewReader(bytes.NewBuffer(data))
	dec, err := ioutil.ReadAll(r)
	fmt.Println(string(dec), err)
}
