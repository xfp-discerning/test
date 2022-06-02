package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

const HelloServiceName = "HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc) //接受的参数为字符串和接口类型
}

//HelloService的私有方法
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	//注册
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	//支持多个TCP链接，然后为每个TCP链接提供RPC服务
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
