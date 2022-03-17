package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for i :=0 ; i < 5; i++{
			fmt.Println("goroutine1")
			if i ==2{
				runtime.Goexit()
			}
		}
	}()
	go func() {
		for i :=0 ; i < 5; i++{
			time.Sleep(time.Second)
			fmt.Println("goroutine2")
		}
	}()


	for i :=0; i< 5; i++{
		//runtime.Gosched()
		fmt.Println("main")
		if i == 2{
			runtime.Goexit()
		}
	}
	
}
