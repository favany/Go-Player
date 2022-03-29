package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 方法一：使用map
	r.GET("/json", func(c *gin.Context) {

		data := map[string]interface{}{
			"name":    "bingo",
			"message": "hello gin!",
			"age":     18,
		}

		// 同样的 gin.H = map[string]interface{}
		data = gin.H{
			"name":    "bingo",
			"message": "hello gin!",
			"age":     18,
		}
		c.JSON(http.StatusOK, data)
	})

	// 方法二：结构体，灵活使用tag来对结构体的字段做定制化操作
	type msg struct {
		Name    string `json:"name"`
		Message string `json:"message"`
		Age     int    `json:"age"`
	}
	r.GET("/struct_json", func(c *gin.Context) {
		data := msg{
			Name:    "Gin",
			Message: "Hello gin, I'm Bingo",
			Age:     18,
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":9000")
}
