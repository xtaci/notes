package main

import "encoding/json"
import "fmt"

func main() {
	x, err := json.Marshal(nil)
	fmt.Println(string(x), err)
}
