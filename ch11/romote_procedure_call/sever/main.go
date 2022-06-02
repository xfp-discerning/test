package main

import (
	"log"
	"net"
	"net/rpc"
)

//RPC服务示例
//RPC规则：方法只能有两个可序列化的参数，其中第二个参数是指针类型，并且返回一个error类型，同时必须是公开的方法。
type HelloSevice struct{}

func (h *HelloSevice) Hello(name string, reply *string) error {
	*reply = "Hello," + name
	return nil
}

func main() {
	//rpc.Register函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数，所有注册的方法会放在“HelloService”服务空间之下
	rpc.RegisterName("HelloSevice", new(HelloSevice))
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error,err:\n", err)
	}
	//rpc.ServeConn函数在该TCP链接上为对方提供RPC服务
	rpc.ServeConn(conn)

}
