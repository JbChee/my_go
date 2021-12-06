package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	consumer,err :=sarama.NewConsumer([]string{"39.97.119.197:2184"},nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	fmt.Println(consumer)

	//指定topic,返回所有分区
	partitions,err :=consumer.Partitions("web_log")
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitions)
	//遍历分区
	for _,partition :=range partitions{
		pc,err :=consumer.ConsumePartition("web_log",int32(partition),sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		fmt.Println(pc)
		defer pc.AsyncClose()

		//省略参数名
		go func(sarama.PartitionConsumer){
			for msg :=range pc.Messages(){
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)

			}
		}(pc)


	}

}
