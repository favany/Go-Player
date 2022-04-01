package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库

	r := gin.Default()
	// 告诉 gin 框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉 gin 框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {

		})
		// 查看所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {

		})

		// 查看某一个代办事项
		v1Group.GET("/todo:id", func(c *gin.Context) {

		})

		// 修改某一个代办事项
		v1Group.PUT("/todo:id", func(c *gin.Context) {

		})

		// 删除某一个代办事项
		v1Group.DELETE("/todo:id", func(c *gin.Context) {

		})
	}

	r.Run(":9001")
}
