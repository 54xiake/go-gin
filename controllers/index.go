package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (i *IndexController) Index(c *gin.Context) {

	//c.JSON(200, gin.H{
	//	"message": "index",
	//})

	//根据完整文件名渲染模板，并传递参数
	c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
		"title": "Main website",
	})
}
