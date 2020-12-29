package controllers

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"go-gin/models"
	"net/http"
)

type IndexController struct {
}

func (i *IndexController) Index(c *gin.Context) {

	fmt.Println("defer begin")
	// 将defer放入延迟调用栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	// 最后一个放入, 位于栈顶, 最先调用
	defer fmt.Println(3)
	fmt.Println("defer end")

	user := &models.User{}
	user.Username = "test"
	user.Password = "123456"

	result, err := user.Create(user)
	if err == nil {
		fmt.Println(result)
		logs.Info(result)
	}

	//c.JSON(200, gin.H{
	//	"message": "index",
	//})

	//根据完整文件名渲染模板，并传递参数
	c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
		"title": "Main website",
	})
}
