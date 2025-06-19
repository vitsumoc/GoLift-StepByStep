package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// query id
	r.GET("/getUserById", func(c *gin.Context) {
		// 获得通过 query 方式传入的查询 id
		id := c.Query("id")

		// 将 id 写入返回值
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("您传入的用户id是：%s", id),
		})
	})

	// query page
	r.GET("userPage", func(c *gin.Context) {
		// 获取通过 query 传入的 page 和 pageSize
		page := c.Query("page")
		pageSize := c.Query("pageSize")

		// 将参数写入响应
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("您传入的页码是：%s，页面大小是：%s", page, pageSize),
		})
	})

	// 通过 param 定位视频 id
	r.GET("/video/:id", func(c *gin.Context) {
		// 获取 path 中的参数
		id := c.Param("id")
		// 写入返回
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("您传入的视频id是：%s", id),
		})
	})

	// 定义接受 JSON 参数的 User 结构体
	type User struct {
		Id       int    `json:"id"`
		Password string `json:"password"`
	}

	// 使用 JSON Body 接收参数
	r.POST("/user/changePasswordByJSON", func(c *gin.Context) {
		// 使用 gin 提供的函数，自动把 body 绑定到 json
		var user User
		c.ShouldBindJSON(&user)
		// 写入返回
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("您传入的用户id是 %d, 用户密码是：%s", user.Id, user.Password),
		})
	})

	// 使用 FormData 接收参数
	r.POST("/user/changePasswordByFormData", func(c *gin.Context) {
		// 接收 formData
		id := c.PostForm("id")
		password := c.PostForm("password")
		// 写入返回
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("您传入的用户id是 %s, 用户密码是：%s", id, password),
		})
	})

	r.Run()
}
