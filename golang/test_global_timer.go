package main

import (
	"sync"
	"time"
)

var _gt global_timer

func init() {
	_gt.init()
	go _gt.generate()
}

type global_timer struct {
	timer_sec chan bool
	timer_min chan bool
	sync.RWMutex
}

func (g *global_timer) init() {
	g.timer_sec = make(chan bool)
	g.timer_min = make(chan bool)
}

func (g *global_timer) get_sec() chan bool {
	g.RLock()
	defer g.RUnlock()
	return g.timer_sec
}

func (g *global_timer) get_min() chan bool {
	g.RLock()
	defer g.RUnlock()
	return g.timer_min
}

func (g *global_timer) generate() {
	t1 := time.After(time.Second)
	t2 := time.After(time.Minute)
	for {
		select {
		case <-t1:
			g.Lock()
			close(g.timer_sec)
			g.timer_sec = make(chan bool)
			g.Unlock()
			t1 = time.After(time.Second)
		case <-t2:
			g.Lock()
			close(g.timer_min)
			g.timer_min = make(chan bool)
			g.Unlock()
			t2 = time.After(time.Minute)
		}
	}
}

func main() {
	const N = 10000
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(x int) {
			timer := _gt.get_sec()
			for {
				select {
				case <-timer:
					println(x)
					timer = _gt.get_sec()
				}
			}
		}(i)
	}
	wg.Wait()
}
