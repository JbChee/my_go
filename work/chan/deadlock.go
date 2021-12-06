package main

import (
	"fmt"
	"time"
)

func main() {


	//eg--1:
	//go func() {
	//	fmt.Println("xxxxxxxxxxxx")
	//	time.Sleep(time.Second*10)
	//}()
	//select {
	//
	//}


	//eg--2:
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	go func() {

		//只有一个sender，直接从sender方关闭即可，否则不能
		defer func() {
			close(ch)
		}()
		for j :=0; j < 10; j++{
			ch <- 1
			time.Sleep(time.Second)
			if j == 9{
				return
			}
		}

	}()

	go func() {
		for v := range ch{
			fmt.Println(v)
		}
	}()

	time.Sleep(time.Second*10)



}
