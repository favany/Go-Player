package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// gorm demo1

// UserInfo --> 数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// 连接 MySQL 数据库
	dsn := "root:@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 创建表 自动迁移（把结构体和数据表进行对应
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	u1 := UserInfo{1, "bingo", "n", "q"}
	db.Create(&u1)

	// 查询
	var u UserInfo
	db.First(&u)
	fmt.Printf("u:%#v\n", u)
	//=>> u:main.UserInfo{ID:0x1, Name:"bingo", Gender:"n", Hobby:"q"}

	// 更新
	db.Model(&u).Update("hobby", "双色球")

	// 删除
	db.Delete(&u)
}
