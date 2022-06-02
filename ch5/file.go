package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// readfilefrombufio()
	// readfilefromioutil()
	fileread()
}
func fileread() {
	fileObj, err := os.Open("./file.go")
	if err == io.EOF {
		fmt.Println("读完了")
		return
	}
	if err != nil {
		fmt.Println("open file failed,error:", err)
		return
	}
	defer fileObj.Close()
	var tmp = make([]byte, 128)
	for {
		n, err := fileObj.Read(tmp[:])
		if err != nil {
			fmt.Println("read from file failed, error:", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:]))
		if n < 128 {
			return
		}
	}
}

//bufio可用于每行读取
func readfilefrombufio() {
	fileObj, err := os.Open("./file.go")
	if err != nil {
		fmt.Println("open file failed,error:", err)
		return
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	//for循环逐行读取文件
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("read from file failed, error:", err)
			return
		}
		fmt.Print(line) //不需要强转
	}
}

//直接读取整个文件
func readfilefromioutil() {
	ret, err := ioutil.ReadFile("./file.go")
	if err != nil {
		fmt.Println("read from file failed, error:", err)
		return
	}
	fmt.Print(string(ret)) //需要用string类型强转一下

}
