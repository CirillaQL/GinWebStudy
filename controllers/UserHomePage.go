package controllers

import (
	"GinWebStudy/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//IndexHTML 函数用于初始化主界面
func IndexHTML(c *gin.Context) {
	//首先验证cookies
	UsernameFromCookie, err := c.Cookie("user")
	if err != nil {
		log.Println("获取cookie失败 Error: ", err)
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
	UsernameFromUrl := c.Query("username")
	if UsernameFromCookie != UsernameFromUrl {
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
	LoadPicture := util.GetPictureList(UsernameFromUrl)

	fmt.Println(LoadPicture)
	c.HTML(http.StatusOK, "photo.html", gin.H{
		"name":  UsernameFromUrl,
		"image": LoadPicture,
	})
}
