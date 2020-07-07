package controllers

import (
	"GinWebStudy/util"
	_ "GinWebStudy/util"
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
	user := util.CheckUser(PhoneNumber)
	ctx.SetCookie("user", user.Name, 3*3600, "/homepage", "localhost", false, true)
	ctx.SetCookie("isLogin", "true", 3*3600, "/homepage", "localhost", false, true)
	ctx.Redirect(http.StatusMovedPermanently, "/homepage?username="+user.Name)
}
