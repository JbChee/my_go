package main

import (
	"fmt"
	"time"
)

//go1结束不影响go2继续运行
//main函数结束，会杀死所有线程

func go1() {
	defer fmt.Println("go1 exit....")
	//defer wg.Done()
	fmt.Println("go1 runing.....")
	go go2()
	time.Sleep(time.Second * 3)
	fmt.Println("return go1....")

}

func go2(){
	defer fmt.Println("go2 exit....")
	//defer wg.Done()
	fmt.Println("go2 runing.....")
	go go3()
	time.Sleep(time.Second * 6)
	fmt.Println("return go2...")
}

func go3(){
	fmt.Println("go3 runing...")
	fmt.Println("return go3....")
}

//var wg sync.WaitGroup
func main() {
	//trace.Start(os.Stderr)
	//defer trace.Stop()
	//wg.Add(2)

	go go1()

	//wg.Wait()

	fmt.Println("main exit......")
}
