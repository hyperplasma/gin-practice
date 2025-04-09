package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个默认路由引擎
	r := gin.Default()
	// 配置路由
	r.GET("/", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello, Hyplus!",
		})
	})
	// 启动HTTP服务，默认为0.0.0.0:8080
	r.Run()
}
