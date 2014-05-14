package main

import (
	"fmt"
)

type X struct {
	A int
}

func main() {
	mm := make([]X, 10)

	temp := X{100}

	for k := range mm {
		mm[k] = temp
	}

	ptr := &mm[5]
	ptr.A = 9999
	mm = append(mm, X{88})
	for k := 0; k < 1024*1024; k++ {
		mm = append(mm, X{88})
	}
	fmt.Println(mm[:10])
}
