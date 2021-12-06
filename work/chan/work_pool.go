package main

import (
	"fmt"
	"time"
)

//另一一种（启动销毁很多go，但是同时运行的只有3个）
var limitCh = make(chan int, 3)
func worker2(workerNum int) {

	for i := 0; i < workerNum; i++ {
		go func(i int) {
			limitCh <- 1
			defer func() {<- limitCh}()
			w(i)

		}(i)
	}
}

func w(i int){
	s := fmt.Sprintf("woker == %v",i)
	fmt.Println( s)
	time.Sleep(time.Second)
}

//工作池（限制最多可以启动 go的数量，性能高一点）
func worker(taskCh <-chan int, n int) {
	N := n
	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				item := <-taskCh
				//do something
				s := fmt.Sprintf("worker %v, data flow = %v", id, item)
				fmt.Println(s)
				time.Sleep(time.Second)
			}
		}(i)
	}
}

func main() {

	//第一种
	//taskCh := make(chan int, 100)
	//go worker(taskCh, 5)
	//
	//for i := 0; i < 200; i++ {
	//	taskCh <- i
	//}
	//for i := 200; i < 10000; i++ {
	//	taskCh <- i
	//}
	//fmt.Println("xxxxxxxxxxxxx------end ")

	//select {
	//case <- time.After(time.Hour):
	//	fmt.Println("xxxxxxxxx")
	//}
	//select {
	//
	//}

	//for {
	//}



	//第二种
	worker2(10000)
	select {
	//default:
	//	fmt.Println("xxxxdeal")
	}

}


