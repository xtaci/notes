package main

import (
	"fmt"
	"sync"
	"time"
)

const ng = 20000
const n = 40
const interval = 150 * time.Millisecond

func main() {
	fmt.Println(ng, "goroutines", n, "iterations")
	testBroadcast()
	testSleep()
}

func testBroadcast() {
	ts := make([][ng]time.Duration, n)

	var started sync.WaitGroup
	started.Add(ng)
	start := make(chan struct{})

	var wait func()

	var wg sync.WaitGroup
	wg.Add(ng)
	for g := 0; g < ng; g++ {
		g := g
		go func() {
			started.Done()
			<-start
			for i := 0; i < n; i++ {
				wait()
				ts[i][g] = time.Since(startTime)
			}
			wg.Done()
		}()
	}
	started.Wait()
	startTime = time.Now()
	broad := newIntervalBroadcaster(interval)
	wait = func() {
		<-broad.WaitChan()
	}
	close(start)
	wg.Wait()
	assess("broadcast", ts)
}

func assess(what string, ts [][ng]time.Duration) {
	var total time.Duration
	nmissed := 0
	for i := range ts {
		for g := 0; g < ng; g++ {
			expect := time.Duration(i+1) * interval
			got := ts[i][g]
			if got-expect > interval {
				nmissed++
			}
			total += got - expect
		}
	}
	fmt.Println(what)
	if nmissed > 0 {
		fmt.Println("  missed", nmissed)
	}
	fmt.Println("  ", total/time.Duration(ng*n))
}

func testSleep() {
	ts := make([][ng]time.Duration, n)

	var started sync.WaitGroup
	started.Add(ng)
	start := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(ng)
	for g := 0; g < ng; g++ {
		g := g
		go func() {
			started.Done()
			<-start
			for i := 0; i < n; i++ {
				sleepWait()
				ts[i][g] = time.Since(startTime)
			}
			wg.Done()
		}()
	}
	started.Wait()
	startTime = time.Now()
	close(start)
	wg.Wait()
	assess("sleep", ts)
}

var startTime time.Time

func sleepWait() {
	now := time.Now()
	d := now.Sub(startTime)
	d = (d + interval) / interval * interval
	sleepDuration := startTime.Add(d).Sub(now)
	time.Sleep(sleepDuration)
}

func afterWait() {
	now := time.Now()
	d := now.Sub(startTime)
	d = (d + interval) / interval * interval
	sleepDuration := startTime.Add(d).Sub(now)
	<-time.After(sleepDuration)
}

func newIntervalBroadcaster(interval time.Duration) *Broadcaster {
	b := NewBroadcaster()
	go func() {
		for {
			sleepWait()
			b.Signal()
		}
	}()
	return b
}

type Broadcaster struct {
	mu sync.RWMutex
	c  chan struct{}
}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{
		c: make(chan struct{}),
	}
}

func (b *Broadcaster) WaitChan() <-chan struct{} {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.c
}

func (b *Broadcaster) Signal() {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.c != nil {
		close(b.c)
	}
	b.c = make(chan struct{})
}
