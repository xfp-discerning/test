package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

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

	//put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "xfp", "handsome")
	cancel()//做一个超时处理
	if err != nil {
		fmt.Println("put to etcd failed, err:", err)
		return
	}

	//get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "xfp")
	cancel()
	if err != nil {
		fmt.Println("get from etcd failed, err:", err)
		return
	}
	for _, kv := range resp.Kvs {
		fmt.Printf("%s : %s\n", kv.Key, kv.Value)
	}
}
