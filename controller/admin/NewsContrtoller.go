package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewsController struct {
}

func (c NewsController) Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "新闻首页")
}

func test() {
	mp := make(map[string]string)
	fmt.Println(mp)
}
