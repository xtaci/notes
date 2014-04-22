package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x, err := json.Marshal([]byte{1, 2, 3, 4, 5})
	fmt.Println(string(x), err)
}
