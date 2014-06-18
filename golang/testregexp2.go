package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`[0-9]+`)
	fmt.Println(re.FindAllString(`{11341234,2,3,4,5}`, -1))
}
