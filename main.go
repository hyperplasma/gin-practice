package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()

	// 配置模板文件
	r.LoadHTMLGlob("templates/**/*")

	// 后台首页
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
			"langs": []string{"Go", "Java", "C", "Python"},
			"newsList": []interface{}{
				&Article{Title: "Go语言基础教程", Desc: "Hello Go!", Content: "content1"},
				&Article{Title: "Java语言基础教程", Desc: "Hello Java!", Content: "content2"},
				&Article{Title: "C语言基础教程", Desc: "Hello C!", Content: "content2"},
			},
			"emptySlice": []int{},
		})
	})

	r.GET("/news", func(c *gin.Context) {
		news := &Article{
			Title:   "Gin框架基础教程",
			Desc:    "Hello Templates",
			Content: "test",
		}
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "I'm data from server",
			"news":  news,
		})
	})

	r.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "This is Goods Page, data from server!",
			"price": 20,
		})
	})

	r.Run()
}
