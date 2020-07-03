package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Get命令时Handler
func RegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

//Post命令时Handler
func RegisterPost(c *gin.Context) {
	username := c.PostForm("name")
	password := c.PostForm("psd")
	//repassword := c.PostForm("repassword")

	//确认密码
	/*
		if password != repassword {
			c.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "密码不一致",
			})
			return
		}*/

	log.Println("用户名：", username, "密码：", password)
	/*
		userRegister := data.User{
			Id:       data.CreateID(),
			Name:     username,
			Password: password,
		}*/

	//检查数据库中是否重复
	/*
		if !data.CheckUser(username) {
			c.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "用户名重复",
			})
		}*/

	//data.SaveToDataBase(userRegister)
	c.Redirect(http.StatusMovedPermanently, "/")
}
