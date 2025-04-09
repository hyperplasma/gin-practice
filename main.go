package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/news", func(c *gin.Context) {
		c.String(http.StatusOK, "This is News page!")
	})

	r.POST("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "POST请求-主要用于增加数据")
	})

	r.PUT("/edit", func(c *gin.Context) {
		c.String(http.StatusOK, "PUT请求-主要用于编辑数据")
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.String(http.StatusOK, "DELETE请求-主要用于删除数据")
	})

	r.Run()
}
