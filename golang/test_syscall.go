package main

import (
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open(os.DevNull)
	if err != nil {
		log.Fatal(err)
	}

	b := make([]byte, 128)
	start := time.Now()
	for i := 0; i < 1000; i++ {
		file.Write(b)
	}
	log.Println(time.Now().Sub(start))
}
