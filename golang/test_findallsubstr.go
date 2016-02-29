package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(findallsubstr("hello"))
	fmt.Println(len(findallsubstr("hello")))
}
func findallsubstr(str string) (strs []string) {
	strs = append(strs, str)
	var runes []rune
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]
		runes = append(runes, r)
	}

	for sz := 1; sz < len(runes); sz++ {
		for i := 0; i+sz <= len(runes); i++ {
			strs = append(strs, string(runes[i:i+sz]))
		}
	}
	return
}
