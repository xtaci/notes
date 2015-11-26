package main

import (
	"fmt"
	"hash/crc32"
	"time"
)

func main() {
	CRC_TABLE := crc32.MakeTable(crc32.Castagnoli)
	start := time.Now()
	data := make([]byte, 10240)
	for i := 0; i < 1000; i++ {
		crc32.Checksum(data, CRC_TABLE)
	}
	fmt.Println(time.Now().Sub(start))
}
