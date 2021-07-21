package main

import (
	"flag"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"go-gin/initialize"
	"net/http"
	"time"
)

var (
	env = flag.String("env", "", "environment dev|prepub|produce")
	//g errgroup.Group

)

func main() {
	initialize.Env = *env
	initialize.InitBasePath()
	initialize.InitConfig()
	initialize.InitDirs()
	initialize.InitLog()
	initialize.InitDB()
	defer initialize.CloseDB()

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

	//性能监控
	//pprof.Register(router)

	ginpprof.Wrap(router)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   90 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

	//server01 := &http.Server{
	//	Addr:":3001",
	//	Handler: router01(),
	//	ReadTimeout:5*time.Second,
	//	WriteTimeout:10*time.Second,
	//}
	//
	//server02 := &http.Server{
	//	Addr:":3002",
	//	Handler:router02(),
	//	ReadTimeout:5*time.Second,
	//	WriteTimeout:10*time.Second,
	//}

	//g.Go(func() error{
	//	return server01.ListenAndServe()
	//})
	//
	//g.Go(func() error{
	//	return server02.ListenAndServe()
	//})
	//
	//if err := g.Wait();err != nil {
	//	logs.Info(err)
	//}

}

func router01() http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK,
			"error": "Welcome 01"})
	})

	return r

}

func router02() http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{"code": http.StatusOK,
				"error": "Welcome server 02"})
	})
	return r
}
