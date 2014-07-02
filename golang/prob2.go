package main

import (
	"fmt"
	//	"html/template"
	//	"log"
	//	"net/url"
	//	"os"
	//	"path/filepath"
	"strconv"
)

type http_client struct {
	buffer []byte
}

type http_protocol struct {
	client *http_client
	buffer []byte
}

type Json struct {
	Status  string
	Message string
}

func (c *http_client) read(buffer []byte) int {
	newbuffer := []byte("12")
	c.buffer = newbuffer
	fmt.Printf("client buffer is:%v\n", c.buffer)

	//buffer = c.buffer
	copy(buffer, c.buffer)
	return len(c.buffer)
}

func (p *http_protocol) read() {
	size := p.client.read(p.buffer)
	fmt.Printf("protocol buffer is:%v\n", p.buffer)
	i, _ := strconv.Atoi(string(p.buffer[:size]))
	fmt.Printf("i is:%v\n", i)

}

func main() {
	test_copy()
}

func test_copy() {
	client := http_client{buffer: make([]byte, 0, 10)}
	protocol := http_protocol{client: &client, buffer: make([]byte, 10, 10)}
	protocol.read()

}
