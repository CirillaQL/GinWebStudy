package controllers

import (
	"GinWebStudy/util"
	_ "GinWebStudy/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//Get命令时Handler
func RegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

//Post命令时Handler
func RegisterPost(c *gin.Context) {
	/*
		从HTML中读取数据
	*/
	username := c.PostForm("name")
	mobile := c.PostForm("mobile")
	password := c.PostForm("psd")
	log.Println("用户注册:    ", "用户名：", username, "手机号码：", mobile, "密码：", password)

	/*
		生成结构体
	*/
	//生成用户
	userRegister := util.User{
		ID:          util.CreateID(),
		Name:        username,
		Password:    password,
		PhoneNumber: mobile,
	}
	userRegister.EncryptPassword()

	ans := util.SaveToDataBase(userRegister)
	util.CreateFileDir(userRegister.Name)
	if ans {
		fmt.Println("注册成功")
		util.CreateFileDir(userRegister.Name)
	} else {
		log.Fatal("插入失败!")
	}

	cookie := util.UserCookie{
		UserName: userRegister.Name,
		Password: userRegister.Password,
		LastTime: 3 * 3600,
	}
	cookie.AddCookieToRedis()

	c.SetCookie("user", userRegister.Name, 3*3600, "/homepage", "localhost", false, true)
	c.SetCookie("password", userRegister.Password, 3*3600, "/homepage", "localhost", false, true)
	c.Redirect(http.StatusMovedPermanently, "/homepage?username="+username)
}
