package main

import (
	"fmt"
	"sync"
	"time"
)

func dosomething(mil time.Duration,wg  *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(mil * time.Millisecond)
	//wg.Add(1)
	fmt.Println("background go")


}

func main() {
	var wg sync.WaitGroup

	wg.Add(4)
	go dosomething(100, &wg)
	go dosomething(120, &wg)
	go dosomething(200, &wg)
	go dosomething(210, &wg)

	wg.Wait()

	fmt.Println("main......")

}
