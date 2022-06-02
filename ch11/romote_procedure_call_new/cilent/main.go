package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const HelloServiceName = "HelloService"

type HelloSeviceClient struct {
	*rpc.Client //一个结构体
}

//没懂
// var HelloSeviceInterface = (*HelloSeviceClient)(nil)

//拨号功能封装，传入协议和地址字段，返回*rpc.Client类型(rpc.Dial的返回类型)和err
func DialHelloService(network, address string) (*HelloSeviceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloSeviceClient{Client: c}, nil
}

//封装client.Call，使之只传入两个字段，且Hello方法是私有方法
func (h *HelloSeviceClient) Hello(request string, reply *string) error {
	return h.Client.Call(HelloServiceName+".Hello", request, reply)
}

func main() {
	client, err := DialHelloService("tcp", "127.0.0.1:20000")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
