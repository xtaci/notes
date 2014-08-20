package main

import (
	"container/list"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

var atExitHooks = list.New()
var atExitHookMutex sync.Mutex
var shouldOsExit = flag.Bool("exit", false, "should we call os.Exit() from a function")
var shouldPanic = flag.Bool("panic", false, "should we panic() out")
var shouldGoPanic = flag.Bool("go-panic", false, "should panic in other go-routine")

func RunAtExitHooks() {
	recoverErr := recover()
	if recoverErr != nil {
		fmt.Fprintf(os.Stderr, "were panicing: %s\n", recoverErr)
	}
	for e := atExitHooks.Back(); e != nil; e = e.Prev() {
		hook := e.Value.(func())
		func(h func()) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Fprintf(os.Stderr, "exit hook failure: %s\n", err)
				}
			}()
			h()
		}(hook)
	}
	if recoverErr != nil {
		panic(recoverErr)
	}
}

func AtExit(h func()) {
	atExitHookMutex.Lock()
	defer atExitHookMutex.Unlock()
	atExitHooks.PushBack(h)
}

func demo1() {
	fmt.Println("Called demo1().")
	AtExit(func() { fmt.Println("Registered in demo1().") })
}

func demo2exit() {
	fmt.Println("Registered in demo2().")
}

func demo2() {
	fmt.Println("Called demo2().")
	AtExit(demo2exit)
}

func demo_exit() {
	fmt.Println("Called demo_exit(), will os.Exit().\n")
	os.Exit(10)
}

func demo_panic() {
	fmt.Println("Raising panic. Let loose the dogs of war.\n")
	panic("oopsydaisy")
}

func demo_elsewhere_panic() {
	fmt.Println("Will panic in thread in half a second.\n")
	go func() {
		time.Sleep(time.Second / 2)
		panic("side-shuffle")
	}()
	time.Sleep(time.Second)
}

func main() {
	flag.Parse()
	if *shouldOsExit && *shouldPanic {
		fmt.Fprintf(os.Stderr, "Can't both exit and panic")
		os.Exit(1)
	}

	fmt.Println("Started.")
	defer RunAtExitHooks()

	AtExit(func() { fmt.Println("Registered in main().") })
	demo1()
	demo2()
	switch {
	case *shouldOsExit:
		demo_exit()
	case *shouldPanic:
		demo_panic()
	case *shouldGoPanic:
		demo_elsewhere_panic()
	}

	fmt.Println("Last line of main().")
}
