package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

func worker(exitchan chan struct{}){
	LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select{
		case <-exitchan:

			break LOOP
		default:
		}

	}
	wg.Done()


}


func main() {
	//有一种特殊的struct{}类型的channel，它不能被写入任何数据，只有通过close()函数进行关闭操作，才能进行输出操作。
	var exitchan = make(chan struct{})
	wg.Add(1)
	go worker(exitchan)
	time.Sleep(time.Second*3)

	exitchan  <- struct{}{}

	close(exitchan)


	wg.Wait()
	fmt.Println("over")

}
