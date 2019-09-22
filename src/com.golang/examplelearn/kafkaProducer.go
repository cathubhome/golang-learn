package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

/**
kafka生产者示例:同步发送消息
参考：
https://www.cnblogs.com/develop-SZT/p/10344589.html
https://www.jianshu.com/p/3d4655cd7054

----小记：有毒,不要踩坑了,fuck,当时以为是go包被墙了没有下载完全,耽误了个把小时
192.168.17.18的主机名为hadoop-senior-1.cathome.com，如果SwitchHosts不开，那么在执行client.SendMessage时会报：
error: send message failed, dial tcp 198.58.118.167:9092: i/o timeout
-----------------------------------------------------------
C:\Users\10157>ping hadoop-senior-1.cathome.com
正在 Ping hadoop-senior-1.cathome.com [198.58.118.167] 具有 32 字节的数据:
来自 198.58.118.167 的回复: 字节=32 时间=277ms TTL=51
来自 198.58.118.167 的回复: 字节=32 时间=278ms TTL=51
来自 198.58.118.167 的回复: 字节=32 时间=277ms TTL=51

198.58.118.167 的 Ping 统计信息:
    数据包: 已发送 = 3，已接收 = 3，丢失 = 0 (0% 丢失)，
往返行程的估计时间(以毫秒为单位):
    最短 = 277ms，最长 = 278ms，平均 = 277ms
*/
func main() {

	config := sarama.NewConfig()
	//等待服务器所有副本保存成功后响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机分区器
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功与失败后的响应
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is a good test, my message is good")

	//使用给定的代理地址与配置创建一个同步生产者
	client, err := sarama.NewSyncProducer([]string{"http://hadoop-senior-1.cathome.com:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}

	defer client.Close()

	//特别注意：本地主机映射要加，否则会报send message failed, dial tcp 198.58.118.167:9092: i/o timeout
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}
	//打印消息分区与偏移量
	fmt.Printf("pid:%v offset:%v\n", pid, offset)

}
