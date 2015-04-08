package main

import "fmt"

type Mutex chan bool

func (m Mutex) Lock()   { m <- true }
func (m Mutex) Unlock() { <-m }

func Exclusive(m Mutex, i *int) {
	m.Lock()
	defer m.Unlock()
	*i++
}

func main() {
	lock := make(Mutex, 1)
	val := 0

	N := 100000
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
