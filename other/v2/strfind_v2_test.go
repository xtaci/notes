package main

import (
	"bytes"
	"io"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"testing"
	"time"
)

func init() {
	go http.ListenAndServe(":6060", nil)
}

type dummyReader struct {
	count int
	max   int
	rnd   *rand.Rand
}

func (dr *dummyReader) Read(p []byte) (n int, err error) {
	var alpha = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ             ")
	if dr.count == dr.max {
		return 0, io.EOF
	}

	remain := len(p)
	idx := 0
	for remain > 0 {
		p[idx] = alpha[dr.rnd.Intn(len(alpha))]
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
	dr.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	return dr
}

func TestFindUniqueString(t *testing.T) {
	t0 := bytes.NewBufferString("   ")
	findUnique(t0, 10)
	t1 := bytes.NewBufferString("a a b b b c")
	findUnique(t1, 10)
	t2 := bytes.NewBufferString("a a a a a a")
	findUnique(t2, 10)
	t3 := bytes.NewBufferString("a b c d e a")
	findUnique(t3, 10)
	t4 := bytes.NewBufferString("a a a a a b")
	findUnique(t4, 10)
}

func TestFindUnique100M(t *testing.T) {
	dr := newDummyReader(100 * 1024 * 1024)
	findUnique(dr, 10*1024*1024)
}
func TestFindUnique1G(t *testing.T) {
	dr := newDummyReader(1024 * 1024 * 1024)
	findUnique(dr, 100*1024*1024)
}

func TestFindUnique10G(t *testing.T) {
	dr := newDummyReader(10 * 1024 * 1024 * 1024)
	findUnique(dr, 1024*1024*1024)
}

func TestFindUnique100G(t *testing.T) {
	dr := newDummyReader(100 * 1024 * 1024 * 1024)
	findUnique(dr, 10*1024*1024*1024)
}
