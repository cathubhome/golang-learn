package main

import (
	"com.golang/logAgent/tailf"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	etcd_client "go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"golang.org/x/net/context"
	"strings"
	"time"
)

//引入etcd的背景：作为日志收集的客户端，由于各个客户端要收集的服务器应用日志的配置都不同，需要对collect的日志路径与推送主题等灵活配置，
//并且要时刻监听收集日志配置的变化采取相应的动作，如某个微服务上线需要对日志收集，这时候监听到ectd键值对的变化，程序会新起个tailf收集与
//生产消息推送到kafka

//etcd要时刻监听配置的变化,不能退出
type EtcdClient struct {
	client *etcd_client.Client

	etcdKeys []string
}

//关于etcd的全局变量
var (
	client *EtcdClient
)

//连接etcd并监控键值对存储的变化
func initEtcd(addr string, etcdKeyPrefix string) (collectConf []tailf.CollectConf, err error) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   strings.Split(addr, ","),
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		logs.Error("connect failed, err:", err)
		return
	}
	client = &EtcdClient{
		client: cli,
	}

	//弥补配置的不规范
	if !strings.HasSuffix(etcdKeyPrefix, "/") {
		etcdKeyPrefix = etcdKeyPrefix + "/"
	}

	for _, ip := range localIpArray {
		//etcd的键命名规则以log.conf中etcd::configKey+服务器ip，因为收集的日志配置不再使用配置文件，而是放到etcd中，每个服务器要收集的日志配置不一定相同
		etcdKey := fmt.Sprintf("%s%s", etcdKeyPrefix, ip)
		logs.Debug("etcd key:%s", etcdKey)
		client.etcdKeys = append(client.etcdKeys, etcdKey)
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
		response, err := client.client.Get(ctx, etcdKey)
		if err != nil {
			continue
		}
		cancelFunc()
		logs.Debug("read from etcd:%v", response.Kvs)
		for k, v := range response.Kvs {
			fmt.Println("etcd key", k, "value:", v)
			if string(v.Key) == etcdKey {
				err = json.Unmarshal(v.Value, &collectConf)
				if err != nil {
					logs.Error("unmarshal failed, err:%v", err)
					continue
				}
				logs.Debug("log config is %v", collectConf)
			}
		}
	}
	fmt.Println("connect etcd success,collectConf:", collectConf)

	initEtcdWatcher(client.etcdKeys)
	return
}

//为每个要监听的键值对单独起一个goroutine
func initEtcdWatcher(etcdKeys []string) {
	for _, key := range etcdKeys {
		go watchKey(key)
	}
}

//监听etcd存储变化
func watchKey(key string) {
	logs.Debug("begin watch key:%s", key)
	//无限循环
	for {
		watchChan := client.client.Watch(context.Background(), key)
		//最新的日志收集配置
		var collectConf []tailf.CollectConf
		//获取etcd上日志配置成功与否的标识
		var getConfSuccess = true
		for watchResponse := range watchChan {
			for _, ev := range watchResponse.Events {
				//若不再收集某服务器的任何日志,则删除该服务器在etcd的日志收集配置的键值对（即cli.Delete(context.Background(),etcdKey)）
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s] 's config deleted", key)
					collectConf = nil
					continue
				}
				//配置发生变更
				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err := json.Unmarshal(ev.Kv.Value, &collectConf)
					if err != nil {
						logs.Error("key [%s] Unmarshal[%s] error:%v ", err)
						getConfSuccess = false
						continue
					}
				}
				logs.Debug("get config from etcd,type:%s key:%q value:%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
			//这个判断不容忽视，情景：由于运营操作不当等因素，etcd上日志配置发生变更被监听到却反序列化失败了，
			// 这时collectConf为空，那么之前上线的系统日志就都没法收集
			if getConfSuccess {
				logs.Debug("get config from etcd success, %v", collectConf)
				tailf.UpdateConfig(collectConf)
			}
		}

	}
}
