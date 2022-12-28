package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 两个包一个都不能少
// 预处理
// 普通处理
// 1. 普通sql语句执行过程
// 1.客户端对sql语句进行占位符替换得到完整的sql语句。
// 2.客户端发送完整sql语句到mysql服务端。
// 3.mysql服务端执行完整的sql语句并将结果返回给客户端。
//
// 预处理执行过程
// 1.将sql语句分为两部分，命令与数据。
// 2.先把命令部分发送给MySQL服务端，MySQL服务端进行sql预处理。
// 3.然后把数据部分发送给MySQL服务端，MySQL服务端对sql语句进行占位符替换。
// 4.mysql服务端执行完整的sql语句并将结果返回给客户端。
func main() {
	//1，获取连接
	db, err := sqlx.Open("mysql", "root:qwe456654.@tcp(localhost:3306)/gopro")
	//2，传递预处理命令
	if err != nil {
		fmt.Println(err)
	}
	sql := "update person set username='jack' where id=?;"
	sql2 := "delete from person  where id=?;"
	prepare, err := db.Prepare(sql)
	fmt.Println(err)
	//3，传递约束
	result, _ := prepare.Exec(1)
	//4，处理获得数据
	fmt.Println(result.RowsAffected())

	//重复再次调用语句
	prepare, _ = db.Prepare(sql2)
	res, _ := prepare.Exec(2)
	fmt.Println(res.RowsAffected())

	//5，关闭连接
	db.Close()

}
