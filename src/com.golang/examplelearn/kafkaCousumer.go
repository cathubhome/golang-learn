package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"sync"
)

/**
kafka消费者样例：
1)等待一组goroutine结束
2)使用add方法设置等待的数量加1
3）使用done方法设置等待的数量减1
4）当等待的数量等于0时，wait函数返回
sync.WaitGroup的使用,有些java线程的程序计数器CountDownLatch的味道哦
*/

var (
	wg sync.WaitGroup
)

func main() {

	consumer, err := sarama.NewConsumer(strings.Split("192.168.17.18:9092", ","), nil)
	if err != nil {
		fmt.Printf("fail to start consumer: %s\n", err)
		return
	}
	partitionList, err := consumer.Partitions("nginx_log")
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("nginx_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		wg.Add(1)
		defer pc.AsyncClose()
		//为每个分区起一个goroutine消费消息
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
			wg.Done()
		}(pc)

	}
	wg.Wait()
	consumer.Close()

}
