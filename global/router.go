package global

import (
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"go-gin/controllers"
)

func InitRouter(router *gin.Engine) {
	controller := controllers.IndexController{}
	router.GET("/", controller.Index)

	router.GET("/ping", func(c *gin.Context) {
		logs.Info("ping")
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	v1 := router.Group("/v1")
	{
		controller := controllers.UserController{}
		v1.GET("/login", controller.Login)
	}

	v2 := router.Group("/v2")
	{
		controller := controllers.UserController{}
		v2.GET("/login", controller.Login)
	}
}
