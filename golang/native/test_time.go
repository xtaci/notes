package main

import (
	"log"
	"time"
)

func main() {
	log.Println(time.Time{}.IsZero())
}
