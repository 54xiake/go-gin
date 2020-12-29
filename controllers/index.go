package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/models"
	"net/http"
)

type IndexController struct {
}

func (i *IndexController) Index(c *gin.Context) {

	user := &models.User{}
	user.Username = "test"
	user.Password = "123456"

	result, err := user.Create(user)
	if err == nil {
		fmt.Println(result)
	}

	//c.JSON(200, gin.H{
	//	"message": "index",
	//})

	//根据完整文件名渲染模板，并传递参数
	c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
		"title": "Main website",
	})
}
