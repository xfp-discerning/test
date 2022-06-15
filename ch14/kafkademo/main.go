package main

import (
	"fmt"

	"github.com/shopify/sarama"
)

//基于sarama第三方库的kafka client
func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送数据需要leader和follower都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出来一个partition
	config.Producer.Return.Successes = true                   //成功交付信息将在success channel返回

	//构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")

	//连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed,err:", err)
		return
	}

	//注册一个关闭
	defer client.Close()
	//发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed, err:", err)
		return
	}
	fmt.Printf("pid:%v,offset:%v\n", pid, offset)
}
