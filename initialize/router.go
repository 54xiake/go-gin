package initialize

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"go-gin/controllers"
	"net/http"
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

	router.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello,"+name)
	})

	//routerouter.GET("/users/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	message := name + " is " + action
	//	c.String(http.StatusOK, message)
	//})

	router.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Guest") //可设置默认值
		//nickname := c.Query("nickname") // 是 c.Request.URL.Query().Get("nickname") 的简写
		c.String(http.StatusOK, fmt.Sprintf("Hello %s ", name))
	})

	userController := controllers.UserController{}
	router.POST("/loginJSON", userController.LoginJson)
	router.POST("/loginForm", userController.LoginForm)
	router.POST("/login", userController.Login)
	router.GET("/uri/:users/:password", userController.LoginURI)
	router.GET("/home", userController.Home)

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

	//JSON/XML/YAML渲染
	router.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	router.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"users"`
			Message string
			Number  int
		}
		msg.Name = "test"
		msg.Message = "hey"
		msg.Number = 123
		// 注意 msg.Name 变成了 "users" 字段
		// 以下方式都会输出 :   {"users": "test", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	router.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"users": "test", "message": "hey", "status": http.StatusOK})
	})

	router.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	//router.GET("/someProtoBuf", func(c *gin.Context) {
	//	reps := []int64{int64(1), int64(2)}
	//	label := "test"
	//	// The specific definition of protobuf is written in the testdata/protoexample file.
	//	data := &protoexample.Test{
	//		Label: &label,
	//		Reps:  reps,
	//	}
	//	// Note that data becomes binary data in the response
	//	// Will output protoexample.Test protobuf serialized data
	//	c.ProtoBuf(http.StatusOK, data)
	//})

	testController := controllers.TestController{}
	router.GET("/test", testController.Test)

	fileController := controllers.FileController{}
	router.POST("/upload", fileController.Upload)
}
