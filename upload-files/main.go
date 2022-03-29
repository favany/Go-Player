package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("./index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("./upload", func(c *gin.Context) {
		// 从请求中读取文件
		f, err := c.FormFile("f1") // 从请求中获取携带的参数是一样的

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else { // 将读取到的文件保存在本地（服务端本地）

			//dst := fmt.Sprintf("./%s", f.Filename) //都可以
			fileName := "老婆我爱你" + f.Filename
			dst := path.Join("./", fileName)
			c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})

		}
	})
	r.Run(":80")
}
