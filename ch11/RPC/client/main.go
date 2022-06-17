package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	//rpc.Dial拨号RPC服务
	client, err := rpc.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		log.Fatal("Dial failed,err:\n", err)
	}
	var reply string
	//在调用client.Call时，第一个参数是用点号链接的RPC服务名字和方法名字，第二和第三个参数分别我们定义RPC方法的两个参数。
	err = client.Call("HelloSevice.Hello", "world", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
