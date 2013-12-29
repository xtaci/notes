package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x, err := json.Marshal(nil)
	fmt.Println(string(x),err)
}
