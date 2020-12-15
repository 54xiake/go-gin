package controllers

import (
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (c *IndexController) Index(ct *gin.Context) {

	ct.JSON(200, gin.H{
		"message": "index",
	})
}
