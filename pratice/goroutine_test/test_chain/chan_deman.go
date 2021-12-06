package main
import (
	"fmt"
	"time"
)

func test1(){
	ch1 :=make(chan int)
	ch2 :=make(chan int)

	//开启线程  send
	go func(){
		for i:=0;i<100;i++{
			ch1 <-i

		}
		close(ch1)
	}()

	//rece
	go func(){
		for{
			i, ok:= <- ch1
			if !ok{
				fmt.Println(" !ok")
				break
			}
			ch2 <- i*i
			// fmt.Println("ch2",ch2)
		}
		
		close(ch2)
	}()

	for i := range ch2{
		fmt.Println("for ch2",i)
	}



}


func test2(){
	//data3 := 300
	var data3 int64
	for i:=0; i<10000;i++{
		go func() {
			data3 ++
			//fmt.Printf(" data1 %v\n",data3)
			//time.Sleep(2*time.Second)
		}()
	}

	//time.Sleep(2*time.Second)
	for i:=0; i<10000;i++ {
		go func() {
			//time.Sleep(1 * time.Second)
			data3 = data3 + 1
			//fmt.Printf("data2 %v\n", data3)
			//time.Sleep(2*time.Second)
		}()
	}

	fmt.Printf("main %v\n",data3)
	time.Sleep(10*time.Second)
	fmt.Printf("main1 %v\n",data3)
}
func main() {
	//test1()
	test2()
	time.Sleep(2*time.Second)
}