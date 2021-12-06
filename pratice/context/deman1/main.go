package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
func worker(){

	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	wg.Done()

}
func main() {
	wg.Add(1)
	go worker()
	wg.Wait()
	fmt.Println("over")

}
