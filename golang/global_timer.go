package main

import (
	"sync"
	"time"
)

var _gt global_timer

func init() {
	_gt.init()
	go func() {
		for {
			<-time.After(time.Minute)
			_gt.signal()
		}
	}()
}

type global_timer struct {
	c chan bool
	sync.RWMutex
}

func (g *global_timer) init() {
	g.c = make(chan bool)
}

func (g *global_timer) waitchan() chan bool {
	g.RLock()
	defer g.RUnlock()
	return g.c
}

func (g *global_timer) signal() {
	g.Lock()
	defer g.Unlock()
	close(g.c)
	g.c = make(chan bool)
}
