package main

import (
	"fmt"
	"sync"
	"time"
)

type Container struct {
	sync.Mutex // <-- Added a mutex
	//counters map[string]int
	counters int
}

//传指针，不然每次调用都是新的container,每份container都持有各自都lock，所以并发都时候，锁没有生效
func (c *Container) inc(name string) {
	c.Lock() // <-- Added locking of the mutex
	defer c.Unlock()
	//c.counters[name]++
	c.counters++
}

func main() {
	//c := Container{counters: map[string]int{"a": 0, "b": 0}}
	c := Container{counters: 0}

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
	}

	go doIncrement("a", 100000)
	go doIncrement("a", 100000)
	//go doIncrement("a", 100000) //无法运行

	// Wait a bit for the goroutines to finish
	time.Sleep(300 * time.Millisecond)
	fmt.Println(c.counters)
}
