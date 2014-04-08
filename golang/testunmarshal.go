package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	Test int32
}

func main() {
	x := `{
			"test":"1"
		}`

	t := &T{}
	err := json.Unmarshal([]byte(x), t)
	fmt.Println(err, t)
}
