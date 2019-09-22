package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {

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

	//1秒内没有设置键值对成功则报错，可以将秒修改为纳秒级别，这样可以测出
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/logagent/conf/", "sample_value")
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/logagent/conf/")
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}

	//监控etcd的存储的键值对变化
	for {
		watch := cli.Watch(context.Background(), "/logagent/conf/")
		for item := range watch {
			for _, ev := range item.Events {
				fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}

	}

}
