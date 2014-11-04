package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.SplitHostPort("192.168.0.1:8888"))
}
