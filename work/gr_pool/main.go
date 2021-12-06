package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)


var (
	p = NewWorkPool(64)
)

// 工作的池子
type WorkPool struct {
	jobList []Job      // 队列
	p       chan Job   // 阻塞式
	cap     int        // 容量大小
	mu      sync.Mutex // 🔐
	once    int32      // 只启动一次有效
}

type Job struct {
	notify chan struct{}                                    // 用于通知， 工作已经完成了
	params []interface{}                                    // 可变参数
	o      func(notify chan struct{}, param ...interface{}) // 工作任务
}

// 工作线程池
func NewWorkPool(cap int) *WorkPool {
	return &WorkPool{cap: cap, p: make(chan Job, cap*4), jobList: make([]Job, 0, cap*4)}
}

// 启动运行
func (c *WorkPool) Run() {
	// 这个函数只有一次运行时有效
	if atomic.CompareAndSwapInt32(&(c.once), 0, 1) {
		go c.h()
		for i := 0; i < c.cap; i++ {
			go c.dispatchJob()
		}
	}

}

func (c *WorkPool) dispatchJob() {
	for v := range c.p {
		v.o(v.notify, v.params...)
	}
}

func (c *WorkPool) Push(j Job) {
	c.mu.Lock()
	c.jobList = append(c.jobList, j)
	c.mu.Unlock()
}

func (c *WorkPool) Block(j Job) {
	c.mu.Lock()
	c.p <- j
	c.mu.Unlock()
}

func (c *WorkPool) h() {
	// 将 jobList 上的任务搬到 p
	for {
		c.mu.Lock()
		//fmt.Println(" ", len(c.p), ", cap:",cap(c.p))
		if len(c.jobList) > 0 {
			if len(c.p) < cap(c.p) {
				head := c.jobList[0]
				c.jobList = c.jobList[1:]
				c.p <- head
			} else {
				fmt.Println("busy", len(c.p))
			}
		}
		c.mu.Unlock()
		time.Sleep(time.Millisecond)
	}
}


func main() {
	go p.Run() // 这个执行，
	//go p.Run() // no work
	//go p.Run() // no work

	for i := 0; i < 4096; i++ {
		go func(idx int) {
			j := Job{notify: make(chan struct{}), o: func(notify chan struct{}, params ...interface{}) {
				defer close(notify)
				fmt.Println("i do a job", params)
				time.Sleep(time.Microsecond*500*2)
			}, params: []interface{}{"hello", "job", "worker", "dataGap", idx}}
			p.Push(j)
			<-j.notify
			fmt.Printf("%v\n",runtime.GOMAXPROCS(8))
			fmt.Printf("you done the job %v \n",idx)
		}(i)
	}
	select {}

}



