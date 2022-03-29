package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 跳转到网站，地址变化
	r.GET("/StatusMovedPermanently", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.liwenzhou.com")
	})

	// 转发，地址不变
	r.GET("/changePath", func(c *gin.Context) {
		c.Request.URL.Path = "/destination" // 把请求的 URI 修改
		r.HandleContext(c)                  // 继续后续的处理
	})

	r.GET("/destination", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello we are at the destination.",
		})
	})

	// 路由重定向

	r.Run(":8080")
}
