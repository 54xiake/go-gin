package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"go-gin/initialize"
	"net/http"
	"time"
)

var (
	env = flag.String("env", "", "environment dev|prepub|produce")
)

func main() {
	initialize.Env = *env
	initialize.InitBasePath()
	initialize.InitConfig()
	initialize.InitDirs()
	initialize.InitLog()
	initialize.InitDB()
	//defer initialize.CloseDB()

	if *env == "produce" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 默认启动方式，包含 Logger、Recovery 中间件
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.Use(initialize.CustomRouterMiddle1)
	router.Use(initialize.LoggerToFile())

	router.LoadHTMLGlob("templates/**/*")
	//router.LoadHTMLGlob("templates/*")

	initialize.InitRouter(router)
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
