package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
var exit string

func worker(){
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)

		if exit == "break"{
			break
		}
	}
	wg.Done()
}


func main() {
	wg.Add(1)
	go worker()

	time.Sleep(time.Second*3)
	exit = "break"
	wg.Wait()
	fmt.Println("over")


}
