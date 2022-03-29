package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:0)/db?charset=utf8mb4&parseTime=True"
	// 注意⚠️ 这里不要使用 := ，我们是给全局变量赋值，然后在 main 函数中使用全局变量 db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库连接（校验 dsn 是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	//db, err := sql.Open("mysql", "root:123456@/db")
	//if err != nil {
	//	panic(err)
	//}
	//print(db)
	//// See "Important settings" section.
	//db.SetConnMaxLifetime(time.Minute * 3)
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(10)
	//
	//fmt.Printf("db: %v\n", db)

	err := initDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("连接成功！")
	}

}

//import (
//	"database/sql"
//	"time"
//
//	_ "github.com/go-sql-driver/mysql"
//)
//
//// ...
//
//db, err := sql.Open("mysql", "user:password@/dbname")
//if err != nil {
//	panic(err)
//}
//// See "Important settings" section.
//db.SetConnMaxLifetime(time.Minute * 3)
//db.SetMaxOpenConns(10)
//db.SetMaxIdleConns(10)
//
//
