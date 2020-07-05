package controllers

import (
	_ "GinWebStudy/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Get Login页面
func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

//Post Login页面
func LoginPost(ctx *gin.Context) {
	PhoneNumber := ctx.PostForm("mobile")
	Password := ctx.PostForm("psd")
	fmt.Println("账号: ", PhoneNumber, "   密码: ", Password)
}
