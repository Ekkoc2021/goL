package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	//1.获取连接:sqlx.Open("驱动名称",账号:密码@tcp(ip:端口)/数据库名称)
	database, _ := sqlx.Open("mysql",
		"root:qwe456654.@tcp(localhost:3306)/gopro")

	var id int
	var name string
	var gender string
	var email string

	//2,查询结果
	rows, _ := database.Query("select * from person;")
	for rows.Next() {
		rows.Scan(&id, &name, &gender, &email)
		fmt.Println(id, name, gender, email)
	}

	row := database.QueryRow("select * from person;") //只获得了一条数据
	row.Scan(&id, &name, &gender, &email)
	fmt.Println(id, name, gender, email)

	//Select(dest interface{}, query string, args ...interface{}) error
	//select方法会将查询到的结果赋值给接口中的实例
	type Person struct { //一定要大写
		//id     int    `db:"id"`
		//name   string `db:"username"`
		//gender string `db:"gender"`
		//email  string `db:"email"`
		Id       int    `db:"id"`
		Username string `db:"username"`
		Gender   string `db:"gender"`
		Email    string `db:"email"`
	}
	var person []Person
	database.Select(&person, "select * from person ;")
	fmt.Println(person)

	//3,关闭
	rows.Close() //close的说明是如果rows中的值已经完全取出会自动关闭
	database.Close()
}

//rows结构体中,基本封装了查询结果集的各种信息,也提供各种方法
//rows封装方法也就寥寥几个,但是功能是相当齐全的
