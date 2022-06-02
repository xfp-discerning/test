package main

import (
	"fmt"
	"net"
)

func procConn(conn net.Conn) {
	//3、与客户端通信
	var tmp [128]byte
	//for循环能多次输入
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed, err:", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
}

func main() {
	//1、本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp server on 12.7.0.0.1:20000 failed , err :", err)
		return
	}
	//2、等待别人来跟我建立连接,一个tcp服务器可以连接多个客户端
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accpet falied, err:", err)
			return
		}
		go procConn(conn)
	}

}
