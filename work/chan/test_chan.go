package main

import (
	. "fmt"
	"time"
)

func test_err() {

	var ch = make(chan int, 10)

	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
			//case v := <-ch:
			//	fmt.Println(v)
			//
		}
	}
	for V := range ch {
		Println(V)
	}
	close(ch)
}

// goroutine 泄漏
func test2_proess(timeout time.Duration) bool {
	//ch := make(chan bool, 1) // 不会泄漏
	ch := make(chan bool)

	go func() {
		time.Sleep((timeout + time.Second))
		ch <- true // 这里会阻塞，因为接受通道关闭⬇️
		println("go exit")
	}()

	select {
	case result := <-ch:
		println("work done")
		return result
	case <-time.After(timeout):
		println("main timeout......")
		return false

	}

}

func test3_wait(timeout time.Duration) bool {
	//ch := make(chan bool, 1)
	ch := make(chan bool)

	go func() {
		time.Sleep((timeout + time.Second))
		ch <- true // 这里panic，因为通道被关闭
		println("go exit")
	}()

	//select {
	//	//case result := <-ch:
	//	//	println("work done")
	//	//	return result
	//	//case <-time.After(timeout):
	//	//	println("main timeout......")
	//	//	return false
	//	//
	//	//}
	defer println("test3_wait exit")
	close(ch)
	time.Sleep(6 * time.Second)
	return <-ch

}

func test4_close(timeout time.Duration) {
	//ch := make(chan bool, 1)
	ch := make(chan int, 5)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		defer println("go exit")
		close(ch)
	}()

	//select {
	//	//case result := <-ch:
	//	//	println("work done")
	//	//	return result
	//	//case <-time.After(timeout):
	//	//	println("main timeout......")
	//	//	return false
	//	//
	//	//}
	defer println("test4_close exit")
	time.Sleep(time.Second)
	println(<-ch) // 即使关闭了ch，ch中还有数据，就可以继续消费数据
	println(<-ch) //
	println(<-ch) //
	println(<-ch) //
	println(<-ch) //
	println(<-ch) //
	println(<-ch) //零值
	println(<-ch) //零值

	return

}

func main() {
	//f, err := os.Create("trace.out")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//
	//err = trace.Start(f)
	//if err != nil {
	//	panic(err)
	//}
	//defer trace.Stop()
	//test()
	//_ = test2_proess(time.Second * 2)
	_ = test3_wait(time.Second * 1)
	//test4_close(time.Second*1)
	//go test_err()

	time.Sleep(time.Second * 3)
}
