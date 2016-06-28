package main

import "fmt"

func _itimediff(later, earlier uint32) int32 {
	return (int32)(later - earlier)
}

func main() {
	fmt.Println(_itimediff(10000, 20000))
	fmt.Println(_itimediff(20000, 10000))
}
