package main

//go连接数据库示例
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

func queryMore(n int) {
	//sql语句
	sqlStr := "select id,name,age from user where id > ?"
	//执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s failed, err: %v\n", sqlStr, err)
		return
	}
	defer rows.Close() //一定要关闭rows//这个scan()函数和单条查询的不一样，里面没有关闭的函数
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scan failed, err: %v\n", err)
			return
		}
		fmt.Printf("u1:%#v\n", u1)
	}

}

//查询单条记录
func queryOne(n int) {
	var u user
	//1、查询语句
	sqlStr := "select id, name, age from user where id=?;"
	//2、执行
	rawObj := db.QueryRow(sqlStr, n) //从数据库连接池中拿一个连接出来去数据库查询单条记录
	//3、拿到结果
	rawObj.Scan(&u.id, &u.name, &u.age) //scan()方法里面有释放连接的代码--->defer r.rows.Close()
	fmt.Printf("user:%#v\n", u)
}

//预处理方式插入,先将sql发往服务器进行编译，一次编译多次执行，大大提升服务器性能
//还可以避免sql注入问题
func prepareInsert() {
	sqlStr := "insert into user(name,age) values(?,?)"
	stmt, err := db.Prepare(sqlStr) //sql在服务端预编译
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	//后续只需要拿stmt进行一些操作
	var m = map[string]int{
		"徐水英": 18,
		"陈望":  23,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
}

//插入
func insert(name string, age int) {
	//sql语句
	sqlStr := "insert into user(name,age) values(?,?)"
	//exec
	ret, err := db.Exec(sqlStr, name, age) //exec执行一次命令（插入，删除，更新）
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	//LastInsertId()能拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("insert get id failed, err:%v\n", err)
		return
	}
	fmt.Println("id:", id)
}

//更新
func updateRow(age int, id int) {
	sqlStr := "update user set age=? where id=?"
	ret, err := db.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("update get id failed, err:%v\n", err)
		return
	}
	fmt.Printf("you have updated %d rows \n", n)
}

//删除
func deleteRows(id int) {
	sqlStr := "delete from user where id=?"
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("delete get id failed, err:%v\n", err)
		return
	}
	fmt.Printf("you have deleted %d rows \n", n)
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

func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("database connected successfully")
	// queryOne(2)
	// queryMore(0)
	// insert("王阳", 23)
	// updateRow(24, 2)
	// deleteRows(4)
	// prepareInsert()
}
