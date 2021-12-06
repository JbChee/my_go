package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1(){

	for i :=1;i<100;i++{
		fmt.Printf("i: %#v",i)
	}
}
func f2(){

	for j :=1;j<100;j++{
		fmt.Printf("j: %#v",j)
	}
}



func main() {
	runtime.GOMAXPROCS(1)
	go f1()
	go f2()
	time.Sleep(time.Second)
	
}