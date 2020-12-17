package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"go-gin/global"
	"net/http"
	"time"
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
	//router.LoadHTMLGlob("templates/*")
	router.LoadHTMLGlob("templates/**/*")

	global.InitRouter(router)
	//router.Run(":8080") //listen and server on 0.0.0.0:8080

	//http.ListenAndServe(":8080", router)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
