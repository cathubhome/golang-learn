package main

import (
	"com.golang/logAgent/kafka"
	"com.golang/logAgent/tailf"
	"fmt"
	"github.com/astaxie/beego/logs"
)

/**
E:\amusement\learning\goProject>go build com.golang\logAgent\main
E:\amusement\learning\goProject>main.exe

借助src/com.golang/examplelearn/logEtcd.go与src/com.golang/examplelearn/kafkaCousumer.go联调
*/
func main() {

	//init config
	fileName := "./config/log.conf"
	err := loadConfig("ini", fileName)
	if err != nil {
		fmt.Printf("load conf failed, err:%v\n", err)
		panic("load conf failed")
		return
	}
	fmt.Printf("loadConfig success\n")

	//init logger
	err = initLogger()
	if err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		panic("initLogger failed")
		return
	}

	fmt.Println("init logger success")

	collectConf, err := initEtcd(appConfig.etcdAddr, appConfig.etcdKey)
	if err != nil {
		fmt.Sprintf("init Etcd error:%v\n", err)
		return
	}
	fmt.Println("init etcd success")

	err = tailf.InitTail(collectConf, appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}
	fmt.Println("init tailf success")

	err = kafka.InitKafka(appConfig.kafkaAddress)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}
	fmt.Println("init kafka success")

	logs.Info("initialize all components success")
	err = serverRun()
	if err != nil {
		logs.Error("server run failed, err:%v", err)
		return
	}
	logs.Info("program exited")
}
