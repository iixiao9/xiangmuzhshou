package main

import (
	"database/sql"
	_ "database/sql/driver"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//定义连接池
var db *sql.DB

//查询
func query() {

}

//插入
func insert() {

}

func initDB() (err error) {
	dsn := "root:123456@tcp(192.168.1.10:3306)/test?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping() //尝试连接mysql
	if err != nil {
		return err
	}
	return nil
}

type user struct {
	id       int
	projects string
	namea    string
	nameb    string
	amount   int
	bmount   int
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed , err %v \n", err)
		return
	}
	fmt.Println("连接数据成功!")
	var u1 user
	sqlStr := `select id, projects, namea, nameb, amount, bmount  from user where id=1;`
	rowObj := db.QueryRow(sqlStr)
	rowObj.Scan(&u1.id, &u1.projects, &u1.namea, &u1.nameb, &u1.amount, &u1.bmount)
	fmt.Printf("u1:%#v\n", u1)
}
