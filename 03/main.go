package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "GET GoLoft!")
	})
	r.POST("/hello", func(c *gin.Context) {
		panic("发生严重故障！")
	})
	r.Run() // 默认运行在 8080 端口
}
