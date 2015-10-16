package rtree

import (
	"github.com/dhconnelly/rtreego"
	"math/rand"
	"testing"
	"time"
)

const (
	KTV = iota
	HOTEL
)

const (
	NEW = iota
	FIVE_YEARS
	TEN_YEARS
	TWENTY_YEARS
)

var tol = 0.01

type Thing struct {
	location rtreego.Point
	name     string
}

func (t *Thing) Bounds() *rtreego.Rect {
	return t.location.ToRect(tol)
}

var (
	rt *rtreego.Rtree
)

func init() {
	rand.Seed(time.Now().Unix())
	rt = rtreego.NewTree(4, 10, 20)
	const N = 10000
	for i := 0; i < N; i++ {
		rt.Insert(&Thing{rtreego.Point{rand.Float64() * N, rand.Float64() * N, HOTEL, float64(i % 4)}, "汉庭"})
	}

	for i := 0; i < N; i++ {
		rt.Insert(&Thing{rtreego.Point{rand.Float64() * N, rand.Float64() * N, KTV, float64(i % 4)}, "天上人间"})
	}
}

func BenchmarkRtree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bb, _ := rtreego.NewRect(rtreego.Point{0, 0, HOTEL, FIVE_YEARS}, []float64{1000, 1000, 0.1, 0.1})
		rt.SearchIntersect(bb)
		bb, _ = rtreego.NewRect(rtreego.Point{0, 0, KTV, TEN_YEARS}, []float64{1000, 1000, 0.1, 0.1})
		rt.SearchIntersect(bb)
	}
}
