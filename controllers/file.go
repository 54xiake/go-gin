package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"go-gin/global"
	"math/rand"
	"net/http"
	"path"
	"strconv"
	"time"
)

type FileController struct {
}

func (f *FileController) Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	logs.Info(file.Filename)

	//上传文件到指定的路径
	rand.Seed(time.Now().UnixNano())
	dst := global.UploadDir + "/" + strconv.Itoa(rand.Intn(1000)) + path.Ext(file.Filename)
	logs.Info("===========" + dst)
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
