package main

// querystring

// GET请求 URL ?后面是querystring参数
// key=value格式，多个key-value用 & 连接
// 例如：http://127.0.0.1:9000/web?query=你好&getQuery=666

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		query := c.Query("query")                                      // 通过 Query 获取请求中携带的querystring参数
		defaultQuery := c.DefaultQuery("defaultQuery", "defaultValue") // 带默认值的查询，取不到就用指定的默认值
		getQuery, ok := c.GetQuery("getQuery")                         // 取到返回 (值, true)；取不到参数，就返回（"", false）
		// 获取浏览器那边发请求携带的 query String 参数
		if !ok {
			// 取不到getQuery
			getQuery = "defaultValue"
		}
		c.JSON(http.StatusOK, gin.H{
			"query":        query,
			"defaultQuery": defaultQuery,
			"getQuery":     getQuery,
		})
	})

	r.Run(":9000")
}
