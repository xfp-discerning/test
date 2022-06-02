package main

import (
	"fmt"
	"net"
	"strings"
)

//udp server端

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	// defer conn.Close()    写在这里不可取，万一并没有建立连接，则关闭不了，出现报错
	if err != nil {
		fmt.Println("fialed,err:", err)
		return
	}
	defer conn.Close() //写到这里更安全
	//不需要建立连接，直接收发数据
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("fialed,err:", err)
			return
		}
		fmt.Println(data[:n])
		reply := strings.ToUpper(string(data[:n]))
		conn.WriteToUDP([]byte(reply), addr)
	}
}
