package main

import (
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	go server()
	time.Sleep(time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	b := make([]byte, 1400)
	if err != nil {
		log.Println(err)
	} else {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			count := 0
			start := time.Now()
			for {
				n, err := conn.Read(b)
				if err != nil {
					log.Println(err)
				}
				count += n
				if count == 1024*1400 {
					break
				}
			}
			log.Println(time.Now().Sub(start))
			wg.Done()
		}()

		for i := 0; i < 1024; i++ {
			_, err := conn.Write(b)
			if err != nil {
				log.Println(err)
			}
		}
		wg.Wait()
	}

}

func server() {
	ln, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buf := make([]byte, 1400)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		_, err = conn.Write(buf[:n])
	}
}
