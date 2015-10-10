package main

import (
	"fmt"
	"github.com/dhconnelly/rtreego"
	"math"
)

const (
	KTV = iota
	HOTEL
)

type Thing struct {
	where *rtreego.Rect
	name  string
}

func (t *Thing) Bounds() *rtreego.Rect {
	return t.where
}

func main() {
	rt := rtreego.NewTree(3, 25, 50)

	p1 := rtreego.Point{1, 1, KTV}
	p2 := rtreego.Point{10, 10, HOTEL}

	r1, _ := rtreego.NewRect(p1, []float64{0.1, 0.1, 0.1})
	r2, _ := rtreego.NewRect(p2, []float64{0.1, 0.1, 0.1})

	rt.Insert(&Thing{r1, "天上人间"})
	rt.Insert(&Thing{r2, "汉庭"})
	inf1 := math.Inf(-1)
	inf2 := math.Inf(1)
	bb, _ := rtreego.NewRect(rtreego.Point{inf1, inf1, HOTEL}, []float64{inf2, inf2, 0.1})
	results := rt.SearchIntersect(bb)
	for k := range results {
		fmt.Println("HOTEL:", results[k])
	}

	bb, _ = rtreego.NewRect(rtreego.Point{inf1, inf1, KTV}, []float64{inf2, inf2, 0.1})
	results = rt.SearchIntersect(bb)
	for k := range results {
		fmt.Println("KTV:", results[k])
	}
}
