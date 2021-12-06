package main

import (
	"fmt"

	"time"
)

func runabc(iterNum int, strNum int) {
	var wait, next, over chan struct{}
	var firstwait chan struct{}
	wait = make(chan struct{})
	firstwait = wait
	fmt.Printf("%#v \n", firstwait)
	fmt.Printf("%#v \n", wait)

	for i := 0; i < strNum; i++ {
		next = make(chan struct{})
		fmt.Printf("next: %#v\n", next)
		over = next
		fmt.Printf("over: %#v\n", over)
		go echoStr(i, wait, next)
		fmt.Printf("over: %#v\n", over)
		wait = next //指针赋值
		fmt.Printf("wait: %#v\n", wait)
	}

	for i := 0; i < iterNum; i++ {
		firstwait <- struct{}{}
		fmt.Printf("firstwait1: %#v\n", firstwait)
		<-over //等待最后一个make出来的 over
	}

	fmt.Printf("wait: %#v\n", wait)
	fmt.Printf("firstwait2: %#v\n", firstwait)

	//close(firstwait)

}

func echoStr(threadNum int, wait chan struct{}, next chan struct{}) {

	str := string('A' + threadNum)

	for _ = range wait {
		fmt.Printf("%d : %s \n", threadNum, str)
		next <- struct{}{}
	}

	//close(next)
	//fmt.Printf("close %d : %s", threadNum, str)

}


//令牌版本
type Token struct{}

func newEchoWorker(i int, curCh chan Token, nextCh chan Token) {
	for {
		token := <-curCh
		fmt.Println(i)
		time.Sleep(time.Second)
		nextCh <- token

	}
}

func echo_str_token(num int) {
	ch := make([]chan Token, num)
	//ch := []chan Token{make(chan Token), make(chan Token),make(chan Token),make(chan Token)}
	for i := 0; i < num; i++ {
		ch[i] = make(chan Token)
	}

	for i := 0; i < num; i++ {
		go newEchoWorker(i, ch[i], ch[(i+1)%num])
	}

	ch[0] <- struct{}{}

	select {}
}
func main() {
	runabc(3,4)
	//echo_str_token(6)
}
