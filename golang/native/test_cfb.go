package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

var IV = []byte{167, 115, 79, 156, 18, 172, 27, 1, 164, 21, 242, 193, 252, 120, 230, 107}
var key = "abcd1234"

func encrypt(block cipher.Block, data []byte) {
	var tbl [aes.BlockSize]byte
	copy(tbl[:], IV)
	block.Encrypt(tbl[:], tbl[:])
	n := len(data) / aes.BlockSize
	for i := 0; i < n; i++ {
		for j := 0; j < aes.BlockSize; j++ {
			data[i*aes.BlockSize+j] = data[i*aes.BlockSize+j] ^ tbl[j]
		}
		copy(tbl[:], data[i*aes.BlockSize:])
		block.Encrypt(tbl[:], tbl[:])
	}

	for j := n * aes.BlockSize; j < len(data); j++ {
		data[j] = data[j] ^ tbl[j%aes.BlockSize]
	}
}
func decrypt(block cipher.Block, data []byte) {
	var tbl [aes.BlockSize]byte
	copy(tbl[:], IV)
	block.Encrypt(tbl[:], tbl[:])
	n := len(data) / aes.BlockSize
	for i := 0; i < n; i++ {
		var next [aes.BlockSize]byte
		copy(next[:], data[i*aes.BlockSize:])
		for j := 0; j < aes.BlockSize; j++ {
			data[i*aes.BlockSize+j] = data[i*aes.BlockSize+j] ^ tbl[j]
		}
		block.Encrypt(next[:], next[:])
		copy(tbl[:], next[:])
	}

	for j := n * aes.BlockSize; j < len(data); j++ {
		data[j] = data[j] ^ tbl[j%aes.BlockSize]
	}
}

func main() {
	plaintext := []byte("By definition of self-synchronising cipher, if part of the ciphertext is lost (e.g. due to transmission errors), then receiver will lose only some part of the original message (garbled content), and should be able to continue correct decryption after processing some amount of input data. This simplest way of using CFB described above is not any more self-synchronizing than other cipher modes like CBC. Only if a whole blocksize of ciphertext is lost both CBC and CFB will synchronize")
	pass := make([]byte, aes.BlockSize)
	copy(pass, []byte(key))
	if block, err := aes.NewCipher(pass); err == nil {
		encrypt(block, plaintext)
		fmt.Println(string(plaintext))
		decrypt(block, plaintext)
		fmt.Println(string(plaintext))
	}
}
