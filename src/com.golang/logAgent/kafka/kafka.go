package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	client sarama.SyncProducer
)

func InitKafka(kafkaAddress string) (err error) {
	config := sarama.NewConfig()
	//等待服务器所有副本保存成功后响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机分区器
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功与失败后的响应
	config.Producer.Return.Successes = true

	//使用给定的代理地址与配置创建一个同步生产者
	client, err = sarama.NewSyncProducer([]string{kafkaAddress}, config)
	if err != nil {
		logs.Error("producer close, error:%v", err)
		return
	}
	logs.Debug("init kafka success")
	return
}

//发送消息到kafka
func SendMsgToKafka(data, topic string) (err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	//partition, offset, err := client.SendMessage(msg)
	_, _, err = client.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed, error:%v data:%v topic:%v", err, data, topic)
		return
	}
	//打印消息分区与偏移量,注释掉，不然读取的日志又写到日志文件中了,推动的日志会不断重复（死循环）
	//logs.Debug("pid:%v offset:%v topic:%v\n", partition, offset,topic)
	return
}
