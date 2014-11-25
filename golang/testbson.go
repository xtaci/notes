package main

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type X struct {
	Data map[int32]int32
	MQ chan int	`bson:"-"`
}

func main() {
	x := X{Data:make(map[int32]int32)}
	x.Data[1] = 1
	buf, err:= bson.Marshal(x)
	fmt.Println(err)
	y := X{}
	bson.Unmarshal(buf, &y)
	fmt.Println(y)
}
