package main

import (
	"fmt"
	"time"
)

var task = make(chan int, 2)

func send() {

	go func() {
		for i := 0; i < 100; i++ {
			select {
			case task <- i:
				s := fmt.Sprintf("xxx ---- %v", i)
				fmt.Println(s)
			default:
				fmt.Println(fmt.Sprintf("default ----- %v", i))

				//使用default ，会跳过需要发送的数据，在阻塞中，执行default分支
			}

		}

	}()

}

func recv() {
	for {
		fmt.Println(<-task)
		time.Sleep(time.Second)
	}
}

func main() {

	//send()
	//go recv()
	ticker := time.NewTicker(time.Second * 5)

	//time.After不能在for循环中使用，会大量创建，但是实际是收到不chan
	endChan := time.After(time.Second * 10)


	for {
		fmt.Println("1s")
		select {
		case  <- endChan:
			fmt.Println("30s end")

		case <- ticker.C:
			fmt.Println("5s")

		}
	}

}
