package main

import "fmt"

type Mutex chan bool

func (m Mutex) Lock()   { <-m }
func (m Mutex) Unlock() { m <- true }

func Exclusive(m Mutex, i *int) {
	m.Lock()
	defer m.Unlock()
	*i++
}

func main() {
	lock := make(Mutex, 1)
	lock <- true
	val := 0

	N := 10
	done := make(chan bool, N)
	for i := 0; i < N; i++ {
		go func() {
			defer func() { done <- true }()
			Exclusive(lock, &val)
		}()
	}
	for i := 0; i < N; i++ {
		<-done
	}

	fmt.Println(val)
}
