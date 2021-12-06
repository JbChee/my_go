package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	//"github.com/gomodule/redigo/redis"
	"strconv"

)

var (
	rdb *redis.Client
)

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "10.100.130.145:6379",
		Password: "123456",  // no password set
		DB:       4,   // use default DB
		PoolSize: 2, // 连接池大小
	})

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//_, err = rdb.Ping(ctx).Result()
	return err
}

func Example() {
	ctx := context.Background()
	if err := initClient(); err != nil {
		return
	}
	tem := make(map[string]int)
	tem["current_water"] = 10
	tem["has_use_water"] = 5

	datas := make([]interface{},0, 4)
	for k, v := range tem{
		datas = append(datas, k, strconv.Itoa(v))
	}
	a := rdb.HMSet(ctx, "test_plam",datas...)
	fmt.Println("a=%v",a)


}


func main(){
	Example()
}


