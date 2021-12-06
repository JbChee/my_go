package main


import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once


func nobufchannel(){
	var a []int
	var b1 chan int
	

	b1 = make(chan int)  //无缓冲
	fmt.Println(a,b1)

	go func(){
		x := <-b1
		fmt.Println( "b1:",x) 
	}()
	b1 <- 10  //hang 住了   无人接受



}

func bufchannel(){
	var b2 chan int
	b2 = make(chan int,10)  //缓冲区

	b2 <- 100

	x := <- b2
	fmt.Println(b2,x)
}


func send(ch1 chan int){
	defer wg.Done()
	for i:=0;i<101;i++{
		ch1 <-i
	}
	// close(ch1)
	once.Do(func() {close(ch1)})  //只执行一次
}


func recv(ch1, ch2 chan int){
	defer wg.Done()
	// for x := range ch1{
	// 	ch2 <- x*x
	// }
	for {
		x, ok := <- ch1
		if !ok{
			break
		}
		ch2 <- x*x
	}
	close(ch2)

}

func channel_test(){
	a := make(chan int,10)
	b := make(chan int,101)
	wg.Add(2)

	go send(a)
	go recv(a,b)
	wg.Wait()
	for ret := range b{
		fmt.Println("ret",ret)
	}



}



func main() {


	nobufchannel()
	bufchannel()
	channel_test()
}