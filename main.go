package main

import (
	"flag"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"go-gin/global"
)

var (
	env = flag.String("env", "", "environment dev|prepub|produce")
)

func main() {
	global.Env = *env
	global.InitBasePath()
	global.InitConfig()
	global.InitDirs()
	global.InitLog()

	if *env == "produce" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 默认启动方式，包含 Logger、Recovery 中间件
	router := gin.Default()
	router.Use(global.LoggerToFile())
	router.GET("/ping", func(c *gin.Context) {
		logs.Info("ping=============")

		global.Logger().Info("记录一下日志=====")

		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	router.Run() //listen and server on 0.0.0.0:8080
}
