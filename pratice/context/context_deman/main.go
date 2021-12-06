package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
func context1(ct context.Context){
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ct.Done():

			break LOOP
		default:
		}
	}
	wg.Done()



}

func context2() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel()都是好的
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	wg.Done()

}


func runcontext2(){
	context2()
}
func main(){
	ct,cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go context1(ct)
	time.Sleep(time.Second*3)
	cancel()
	wg.Wait()
	time.Sleep(time.Second*3)
	fmt.Println("over")

	//runcontext2()


}