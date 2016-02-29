package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {
	c, err := net.DialTimeout("udp", "192.168.2.220:8125", 2*time.Second)
	if err != nil {
		log.Println(err)
		return
	}

	b := []byte("test.test")
	start := time.Now()
	for i := 0; i < 1000; i++ {
		c.Write(b)
	}
	log.Println(time.Now().Sub(start))

	start = time.Now()
	for i := 0; i < 1000; i++ {
		rand.Float32()
	}
	log.Println(time.Now().Sub(start))
}
