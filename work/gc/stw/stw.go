package main

import (
	"runtime"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(time.Second*10)
		}
	}()

	time.Sleep(time.Millisecond)
	runtime.GC()
	println("OK")
}