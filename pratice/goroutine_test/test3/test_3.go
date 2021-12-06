package main

import (
	"fmt"

)

var ch1 chan int  //生命一个传递整型 的通道
var ch2 chan bool
var ch3 chan []int



func f1() {
	var ch chan int
	fmt.Println(ch)
}

func f2(){
	ch4 := make(chan int)
	ch5 := make(chan bool)
	ch6 := make(chan []int)

	fmt.Println(ch4,ch5,ch6)
	ch := make(chan int)
	ch <- 10  //把 10 发送到 ch中

	rece_ch := <- ch
	fmt.Println(rece_ch)
	<- ch

}

//无缓冲通道
func f3(c chan int){
	ret := <-c
	
	fmt.Println("ret",ret)

}

//有缓冲通道
func f4() {
	ch :=make(chan int,1)
	ch <-1000
	fmt.Printf("ch is %#v",ch)
	
}


func main() {
	
	f1()
	// f2()
	ch := make(chan int)
	go f3(ch)

	ch <-10
	fmt.Println("send success")
	f4()

}