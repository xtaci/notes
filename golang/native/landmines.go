package main

import (
	"errors"
	"fmt"
)

func print(pi *int) { fmt.Println(*pi) }

var ErrDidNotWork = errors.New("did not work")

func main() {
	for i := 0; i < 10; i++ {
		//defer fmt.Println(i)
		//defer func() { fmt.Println(i) }()
		//defer func(i int) { fmt.Println(i) }(i)
		//defer print(&i)
		//go fmt.Println(i)
		//	go func() { fmt.Println(i) }()
	}

	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))

}
func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	fmt.Println(reallyDoIt, err)
	return err
}

func tryTheThing() (string, error) {
	return "xxx", errors.New("haha")

}
