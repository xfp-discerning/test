package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init()
)

var db *sql.DB //数据库的连接池

type user struct {
	id   int
	name string
	age  int
}

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/school"
	//     用户名 密码  协议  127.0.0.1是指本机的IP地址,3306是Mysql的端口
	db, err = sql.Open("mysql", dsn) //不会检验用户名和密码的正确性
	//dsn的格式出错时会报错
	if err != nil {
		return
	}
	err = db.Ping() //校验用户名和密码
	if err != nil {
		return
	}
	//设置连接池的最大连接数
	db.SetMaxOpenConns(10)
	//设置最大闲置数
	db.SetConnMaxIdleTime(5)
	return
}

//事务操作
func transactionDemo() {
	//1.开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed, err:%v\n", err)
		return
	}
	//2.执行多个sql
	sqlStr1 := "update user set age=age-2 where id =1"
	sqlStr2 := "update user set age=age+2 where id =3"
	// sqlStr2 := "update xxxx set age=age+2 where id =3" //执行sql2出错,回滚
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		//回滚
		tx.Rollback()
		fmt.Println("执行sql1出错,回滚")
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
		fmt.Println("执行sql2出错,回滚")
		return
	}
	//3.上面执行成功后，提交本次事务
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("提交出错了,回滚")
		return
	}
	fmt.Println("事务执行成功")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("database connected successfully")
	transactionDemo()
}
