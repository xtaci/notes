package main

import (
	"fmt"
	"github.com/juju/testing/checkers"
)

type X struct {
	V int
}

type Y struct {
	x *X
}

func main() {
	a := Y{x: &X{1}}
	b := Y{x: &X{2}}

	fmt.Println(checkers.DeepEqual(a, b))
}
