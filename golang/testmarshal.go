package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x, err := json.Marshal(int64(12341231234134134))
	fmt.Println(string(x), err)
}
