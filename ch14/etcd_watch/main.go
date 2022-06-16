package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)


//设置一个哨兵，监测某个key的变换，包括更新删除等等
func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect etcd failed, err:", err)
		return
	}
	fmt.Println("connect etcd success")
	defer cli.Close()

	//watch
	ch := cli.Watch(context.Background(), "xfp")
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("type:%v, key:%v, value:%v\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
		}
	}
}
