package main

//go连接数据库示例
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init()
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/school"
	//     用户名 密码  协议  127.0.0.1是指本机的IP地址,3306是Mysql的端口
	db, err := sql.Open("mysql", dsn) //不会检验用户名和密码的正确性
	//dsn的格式出错时会报错
	if err != nil {
		panic(err)

	}
	err = db.Ping() //校验用户名和密码
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dsn, err)
		return

	}
	fmt.Println("database connected successfully")
}
