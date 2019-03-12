package main

import (
	"io"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"testing"
)

func init() {
	go http.ListenAndServe(":6060", nil)
}

type dummyReader struct {
	count int
	max   int
}

func (dr *dummyReader) Read(p []byte) (n int, err error) {
	var alpha = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ ")
	if dr.count == dr.max {
		return 0, io.EOF
	}

	remain := len(p)
	idx := 0
	for remain > 0 {
		p[idx] = alpha[rand.Intn(len(alpha))]
		idx++
		remain--
		dr.count++
		if dr.count == dr.max {
			return idx, io.EOF
		}
	}

	return idx, nil
}

func newDummyReader(cap int) *dummyReader {
	dr := new(dummyReader)
	dr.max = cap
	return dr
}

func TestFindUnique100M(t *testing.T) {
	dr := newDummyReader(100 * 1024 * 1024)
	findUnique(dr, 128*1024*1024)
}
func TestFindUnique1G(t *testing.T) {
	dr := newDummyReader(1024 * 1024 * 1024)
	findUnique(dr, 128*1024*1024)
}

func TestFindUnique10G(t *testing.T) {
	dr := newDummyReader(10 * 1024 * 1024 * 1024)
	findUnique(dr, 128*1024*1024)
}

func TestFindUnique100G(t *testing.T) {
	dr := newDummyReader(100 * 1024 * 1024 * 1024)
	findUnique(dr, 128*1024*1024)
}
