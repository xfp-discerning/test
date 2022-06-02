package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//udp客户端

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("连接失败，err:", err)
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	var reply [1024]byte
	for {
		fmt.Println("请输入内容：")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		conn.Write([]byte(msg))
		//收到回复的内容
		n, _, err := conn.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("failed,err:", err)
			return
		}
		fmt.Println("收到回复信息：", reply[:n])
	}
}
