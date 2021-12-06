package main

import (
	"fmt"
	"github.com/Shopify/sarama"

)

func main() {
	//配置
	config :=sarama.NewConfig()
	config.Producer.RequiredAcks =sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	//构造消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a log msg")

	//连接
	client, err :=sarama.NewSyncProducer([]string{"39.97.119.197:9093"},config)
	if err !=nil{
		fmt.Println("connect kafka failed ",err)
		return
	}
	defer client.Close()

	//发送消息
	pid,offset,err :=client.SendMessage(msg)
	if err!=nil{
		fmt.Println("send msg failed ")
		return
	}
	fmt.Printf("pid%v,,offset%v",pid,offset)



}
