package main

import (
	"log"
	"runtime"
)

func PrintPanicStack() {
	i := 0
	funcName, file, line, ok := runtime.Caller(i)
	for ok {
		log.Printf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
		i++
		funcName, file, line, ok = runtime.Caller(i)
	}
}

func main() {
	PrintPanicStack()
}
