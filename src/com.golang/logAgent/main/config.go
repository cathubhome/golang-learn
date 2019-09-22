package main

import (
	"com.golang/logAgent/tailf"
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
)

var (
	appConfig *logAgentConfig
)

/**
工程配置的结构体
*/
type logAgentConfig struct {
	//日志级别
	logLevel string
	//日志路径
	logPath string
	//kafka地址
	kafkaAddress string
	//channel通道大小
	chanSize int
	//收集日志的配置
	collectConfig []tailf.CollectConf

	//etcd配置
	etcdAddr string
	etcdKey  string
}

/**
加载配置
*/
func loadConfig(adapterName, fileName string) (err error) {
	conf, err := config.NewConfig(adapterName, fileName)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}
	appConfig = &logAgentConfig{}
	appConfig.logLevel = conf.String("log::logLevel")
	if len(appConfig.logLevel) == 0 {
		fmt.Println("appConfig.logLevel invalid")
		return
	}
	appConfig.logPath = conf.String("log::logPath")
	if len(appConfig.logPath) == 0 {
		fmt.Printf("appConfig.logPath invalid")
		return
	}

	appConfig.chanSize, err = conf.Int("collect::chanSize")
	if err != nil {
		fmt.Printf("appConfig.chanSize invalid,err:%v\n", err)
		return
	}

	appConfig.kafkaAddress = conf.String("kafka::address")
	if len(appConfig.kafkaAddress) == 0 {
		fmt.Printf("appConfig.kafkaAddress invalid\n")
		return
	}

	appConfig.etcdAddr = conf.String("etcd::addr")
	if len(appConfig.etcdAddr) == 0 {
		fmt.Printf("appConfig.etcd address invalid\n")
		return
	}

	appConfig.etcdKey = conf.String("etcd::configKey")
	if len(appConfig.etcdKey) == 0 {
		fmt.Printf("appConfig.etcd config key invalid\n")
		return
	}

	appConfig.etcdKey = conf.String("etcd::configKey")
	if len(appConfig.etcdKey) == 0 {
		fmt.Printf("appConfig.etcd config key invalid\n")
		return
	}

	err = loadCollectConfig(conf)
	if err != nil {
		fmt.Printf("load collect conf failed, err:%v\n", err)
		return
	}
	return
}

/**
加载收集日志配置-日志文件与kafka主题
*/
func loadCollectConfig(conf config.Configer) (err error) {
	var cc tailf.CollectConf
	cc.LogPath = conf.String("collect::logPath")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::logPath")
		return
	}

	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err = errors.New("invalid collect::topic")
		return
	}

	appConfig.collectConfig = append(appConfig.collectConfig, cc)
	return

}
