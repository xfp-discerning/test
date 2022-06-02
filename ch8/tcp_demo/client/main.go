package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//与server段建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial failed, err: ", err)
		return
	}
	//发送数据
	fmt.Println("please give an input")
	//os.stdin捕捉终端上的输入
	read := bufio.NewReader(os.Stdin)
	for {
		// fmt.Scanln(&msg)
		msg, _ := read.ReadString('\n') //读到换行
		//将字符串前后端的空白都去掉
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}
