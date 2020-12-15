package main

import (
	"flag"
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
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.Use(global.CustomRouterMiddle1)
	router.Use(global.LoggerToFile())
	global.InitRouter(router)
	router.Run() //listen and server on 0.0.0.0:8080
}
