package main

import (
	"fmt"
	"misc/geoip"
	"net"
)

func main() {
	ip := net.IP(make([]byte, 4))

	for a := 0; a <= 255; a++ {
		ip[0] = byte(a)
		for b := 0; b <= 255; b++ {
			ip[1] = byte(b)
			fmt.Println(ip)
			for c := 0; c <= 255; c++ {
				ip[2] = byte(c)
				for d := 0; d <= 255; d++ {
					ip[3] = byte(d)
					geoip.QueryCity(ip)
				}
			}
		}
	}
}
