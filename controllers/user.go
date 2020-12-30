package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-gin/models"
	"net/http"
)

type UserController struct {
}

type Login struct {
	User     string `form:"username" json:"users" uri:"users" xml:"users"  binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

// Example for binding JSON ({"users": "test", "password": "123456"})
func (u *UserController) LoginJson(c *gin.Context) {
	var json Login
	//其实就是将request中的Body中的数据按照JSON格式解析到json变量中
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if json.User != "test" || json.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func (u *UserController) LoginForm(c *gin.Context) {
	var form Login
	//方法一：对于FORM数据直接使用Bind函数, 默认使用使用form格式解析,if c.Bind(&form) == nil
	// 根据请求头中 content-type 自动推断.
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if form.User != "test" || form.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func (u *UserController) Login(c *gin.Context) {
	var form Login
	//方法二: 使用BindWith函数,如果你明确知道数据的类型
	// 你可以显式声明来绑定多媒体表单：
	// c.BindWith(&form, binding.Form)
	// 或者使用自动推断:
	if c.BindWith(&form, binding.Form) == nil {
		if form.User == "users" && form.Password == "password" {
			c.JSON(200, gin.H{"status": "you are logged in ..... "})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	}
}

func (u *UserController) LoginURI(c *gin.Context) {
	var login Login
	if err := c.ShouldBindUri(&login); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"username": login.User, "password": login.Password})
}

func (u *UserController) Home(c *gin.Context) {

	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"title": "Users",
	})
}

func (u *UserController) Create(c *gin.Context) {
	user := &models.User{}
	user.Username = "test"
	user.Password = "123456"

	result, err := user.Create(user)
	if err == nil {
		fmt.Println(result)
		logs.Info(result)
		c.JSON(200, gin.H{"data": result})
	} else {
		c.JSON(500, gin.H{"error": err})
	}

}
