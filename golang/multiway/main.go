package main

import (
	"flag"
)

var tp = flag.String("type", "front", "server or front")

func init() {
	flag.Parse()
}

func main() {
	switch *tp {
	case "server":
		start_multiway_server()
	case "front":
		start_multiway_front()
	}

	go sig_handler()
}
