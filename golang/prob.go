package main

import (
	"fmt"
)

type http_client struct {
	buffer []string
}

type http_protocol struct {
	client *http_client
	buffer []string
}

func (c *http_client) read(buffer []string) {
	newbuffer := []string{"abc"}
	c.buffer = newbuffer
	fmt.Printf("client buffer is:%v\n", c.buffer)

	//buffer = c.buffer
	copy(buffer, c.buffer)
}

func (p *http_protocol) read() {
	p.client.read(p.buffer)
	fmt.Printf("protocol buffer is:%v\n", p.buffer)

}

func main() {
	// s1 := []string{"a"}
	// s2 := []string{"b"}
	// fmt.Printf("s1 is:%v\n", s1)
	// fmt.Printf("s2 is:%v\n", s2)
	// s2 = s1
	// fmt.Printf("s2 is:%v\n", s2)
	// copy(s2, s1)
	// fmt.Printf("s2 is:%v\n", s2)

	client := http_client{buffer: make([]string, 0, 10)}
	protocol := http_protocol{client: &client, buffer: make([]string, 10)}
	protocol.read()
}
