package initialize

import (
	"fmt"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-gin/global"
	"os"
	"path"
	"time"
)

var (
	BasePath string //项目根路径
	TempDir  string //临时目录
	LogDir   string //日志目录
	IniConf  config.Configer
	Env      string //运行环境 dev prepub produce
)

func InitBasePath() {
	BasePath, _ = os.Getwd()
}

func InitConfig() {
	var err error
	IniConf, err = config.NewConfig("ini", BasePath+"/config/main.ini")
	if err != nil {
		os.Exit(1)
	}
}

func InitDirs() {
	tempDir, _ := IniConf.String("tempDir")
	TempDir = BasePath + "/" + tempDir
	if _, err := os.Stat(TempDir); err != nil && os.IsNotExist(err) {
		os.Mkdir(TempDir, 0766)
	}

	logDir, _ := IniConf.String("logDir")
	if logDir == "" {
		logDir = BasePath + "/" + tempDir + "/logs"
	} else {
		logDir = BasePath + "/" + logDir
	}
	if _, err := os.Stat(logDir); err != nil && os.IsNotExist(err) {
		os.Mkdir(logDir, 0766)
	}
	LogDir = logDir

	uploadDir, _ := IniConf.String("uploadDir")
	uploadDir = BasePath + "/" + uploadDir
	if _, err := os.Stat(uploadDir); err != nil && os.IsNotExist(err) {
		os.Mkdir(uploadDir, 0766)
	}
	global.UploadDir = uploadDir
}

func InitLog() {
	if Env == "dev" {
		logs.SetLogger(logs.AdapterConsole)
	}
	logs.SetLogger(logs.AdapterFile, `{"filename":"`+LogDir+`/go-gin.log`+`"}`)

	if Env == "produce" {
		logs.SetLevel(logs.LevelInformational)
	}

	logs.EnableFuncCallDepth(true) //开启行号
	logs.SetLogFuncCallDepth(3)    //调用层级

}

func Logger() *logrus.Logger {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	//logger.SetFormatter(&logrus.TextFormatter{
	//	DisableColors: true,
	//	FullTimestamp: true,
	//	TimestampFormat: "2006-01-02 15:04:05",
	//})
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//logger.SetReportCaller(true)

	return logger
}

func LoggerToFile() gin.HandlerFunc {
	logger := Logger()
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		//日志格式
		logger.Infof("|%3d|%13v|%15s|%s|%s|",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
