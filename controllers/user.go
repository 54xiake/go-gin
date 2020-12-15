package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (c *UserController) Login(*gin.Context) {
	fmt.Println("login")
	return
}
