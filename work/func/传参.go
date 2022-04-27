package main

import (
	"fmt"
	"time"
)

func run(){

	testA := []string{"testA"}
	fmt.Printf("p = %p \n", &testA)

	testIntA := 100
	fmt.Printf("testIntA p = %p \n", &testIntA)
	go func() {
		time.Sleep(time.Second  * 1)
		fmt.Printf("go p = %p \n", &testA)
		fmt.Printf("go testIntA p = %p", &testIntA)
	}()

	testIntA = 200


}

func main(){

	run()
	//aa()

	time.Sleep(time.Second * 10)

}

