package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := float64(-1.0)
	y := float64(1.0)
	fmt.Println(int32(x))
	fmt.Println(int32(y))

	for i := 0; i < 4294967296; i++ {
		tmp, _ := strconv.ParseFloat(fmt.Sprint(i)+".0", 64)
		if int(tmp) != i {
			fmt.Println(tmp, i)
		}
	}
}
