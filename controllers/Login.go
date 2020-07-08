package controllers

import (
	"GinWebStudy/util"
	_ "GinWebStudy/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

	//从数据库中读取用户信息
	user := util.CheckUser(PhoneNumber)
	/*
		密码比较
		错误的话跳转的登录错误的页面，正确的话跳转到主页
	*/
	fmt.Println(user)
	result := user.CheckPassword(Password)
	if result != true {
		ctx.Redirect(http.StatusMovedPermanently, "/loginWrong")
	} else {
		cookie := util.UserCookie{
			UserName: user.Name,
			Password: user.Password,
			LastTime: 3 * 3600,
		}
		cookie.AddCookieToRedis()
		ctx.SetCookie("user", user.Name, 3*3600, "/homepage", "localhost", false, true)
		ctx.SetCookie("isLogin", "true", 3*3600, "/homepage", "localhost", false, true)
		ctx.Redirect(http.StatusMovedPermanently, "/homepage?username="+user.Name)
	}
}

//LoginWrong 页面
func LoginWrong(c *gin.Context) {
	c.HTML(http.StatusOK, "loginWrong.html", gin.H{})
	time.Sleep(3 * time.Millisecond)
	c.Redirect(http.StatusMovedPermanently, "/login")
}
