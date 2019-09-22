package main

import (
	"com.golang/logAgent/kafka"
	"com.golang/logAgent/tailf"
	"github.com/astaxie/beego/logs"
)

func serverRun() (err error) {
	for true {
		msg := tailf.GetMsgLine()
		err = sendMsgToKafka(msg)
		if err != nil {
			logs.Error("send msg to kafka failed, error:%v", err)
			continue
		}
	}
	return
}

func sendMsgToKafka(msg *tailf.TextMsg) (err error) {
	err = kafka.SendMsgToKafka(msg.Msg, msg.Topic)
	return
}
