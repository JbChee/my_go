package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"

)

//声明全局
var (
	pool *redis.Pool
)

func Newpool(add string,pw string) *redis.Pool{
	p := &redis.Pool{
		MaxIdle:3,
		IdleTimeout: 10*time.Second,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp",add,redis.DialPassword(pw))
		},

	}

	return p
}


func main() {
	add := "39.97.119.197:6379"
	pw := "123456"
	pool = Newpool(add,pw)
	fmt.Println("pool",pool)

	c := pool.Get()  //获取连接


	c.Do()

	defer c.Close()

	//c.Do("get")


}

