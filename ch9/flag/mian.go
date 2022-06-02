package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "徐发鹏", "please input your name")
	age := flag.Int("age", 23, "the real age")
	flag.Parse() //解析命令传入的参数，必须要加
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(flag.Args())  //返回命令行参数后面的其他参数，[]string类型
	fmt.Println(flag.NArg())  //返回命令行参数后面的其他参数的个数
	fmt.Println(flag.NFlag()) //返回命令行参数的个数

}
