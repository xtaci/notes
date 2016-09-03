package main

import (
	"encoding/json"
	"log"
)

type Config struct {
	nodelay, interval, resend, nc int
}

/*
func (c *Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.nodelay)
}
*/

func main() {
	bts, _ := json.Marshal(&Config{})
	log.Println(string(bts))
}
