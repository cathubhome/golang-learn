package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

const (
	EtcdKey = "/oldboy/backend/logagent/config/192.168.17.1"
)

type LogConf struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// 修改etcd的日志配置
func SetLogConfToEtcd() {
	//连接Etcd
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"hadoop-senior-1.cathome.com:2379"},
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	defer cli.Close()

	fmt.Println("connect success")

	_, err = cli.Delete(context.Background(), EtcdKey)
	return

	var logConfArr []LogConf
	logConfArr = append(
		logConfArr,
		LogConf{
			Path:  "E:/amusement/learning/goProject/logs/application.log",
			Topic: "nginx_log",
		},
	)
	logConfArr = append(
		logConfArr,
		LogConf{
			Path:  "D:/logs/application/vehicle.2019-04-20.log",
			Topic: "nginx_log",
		},
	)

	data, err := json.Marshal(logConfArr)
	if err != nil {
		fmt.Println("json failed, ", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func main() {
	SetLogConfToEtcd()
}
