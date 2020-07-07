package controllers

import (
	"GinWebStudy/util"
	_ "GinWebStudy/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	//转化密码为密文
	psd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("密码加密失败!  Error:", err)
	}

	//生成用户
	userRegister := util.User{
		ID:          util.CreateID(),
		Name:        username,
		Password:    string(psd),
		PhoneNumber: mobile,
	}

	ans := util.SaveToDataBase(userRegister)
	if ans {

		fmt.Println("注册成功")
	} else {
		log.Fatal("插入失败!")
	}

	c.SetCookie("user", username, 3*3600, "/homepage", "localhost", false, true)
	c.SetCookie("isLogin", "true", 3*3600, "/homepage", "localhost", false, true)
	c.Redirect(http.StatusMovedPermanently, "/homepage?username="+username)
}
