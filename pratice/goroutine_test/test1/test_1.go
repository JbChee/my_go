package main

import (
	"fmt"

	"sync"
)

func hello() {
	fmt.Println("hello world Goroutine")
}

var wg sync.WaitGroup

func more_hello(i int) {
	defer wg.Done()
	fmt.Printf("i is %#v\n",i)
}

func run_() {
	for i :=0;i<100;i++{
		wg.Add(1)
		go more_hello(i)
	}
}

func main() {
	go hello()
	fmt.Println("main goroutine done")
	// time.Sleep(time.Second)

	run_()
	wg.Wait()
}