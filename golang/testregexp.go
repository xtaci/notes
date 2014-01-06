package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`/([a-z]+) ([0-9]+)`)
	fmt.Println(re.FindStringSubmatch("/lv 3"))
	slice := re.FindStringSubmatch("/res 10000")
	fmt.Println(slice[1])
	switch slice[1] {
	case "res":
		fmt.Println(slice)
	}
	fmt.Println(re.FindStringSubmatch("/res  10000 "))
	fmt.Println(re.FindStringSubmatch("res 10000 "))

	reg, _ := regexp.Compile("(?i:fuck)")
	fmt.Println(reg.ReplaceAllString("FUCK YOU", "*"))
	fmt.Println(reg.ReplaceAllString("fuck you", "*"))
}
