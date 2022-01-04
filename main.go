package main

import (
	"database/sql"
	_ "database/sql/driver"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//定义连接池
var db *sql.DB

//单查询
func queryOne(id int) {
	var u1 user
	sqlStr := `select id, projects, namea, nameb, amount, bmount  from user where id=?;`
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.projects, &u1.namea, &u1.nameb, &u1.amount, &u1.bmount)
	fmt.Printf("u1:%#v\n", u1)
}

//多查询
func queryMore(id int) {
	sqlStr := `select id, projects, namea, nameb, amount, bmount  from user where id > ?;`
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("exec %s query failed, error: %v\n", sqlStr, err)
		return
	}
	defer rows.Close()
	//循环读取
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.projects, &u1.namea, &u1.nameb, &u1.amount, &u1.bmount)
		if err != nil {
			fmt.Printf("scan failed ,err %v\n", err)
		}
		fmt.Printf("u1:%#v\n", u1)
	}
}

//插入
func insert() {
	sqlStr := `insert into user(projects, namea, nameb, amount, bmount) values("项目名称","用户1","用户2",100,200)`
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("inster failed,err %v\n", err)
		return
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
	}
	fmt.Println("id", id)
}

//更新数据
func updateRow(Namount int, Nid int) {
	sqlStr := `update user set amount=? where id = ?`
	ret, err := db.Exec(sqlStr, Namount, Nid)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据\n", n)
}

//删除数据
func deleteRow(id int) {
	sqlStr := `delete from user where id=?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete %v failed\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Printf("删除了%d行\n", n)

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
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)
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
	//queryOne(2)
	//insert()
	//updateRow(99, 4)

	queryMore(0)
	deleteRow(3)
}
