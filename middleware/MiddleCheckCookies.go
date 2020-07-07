package middleware

import (
	"GinWebStudy/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//登录检查中间件
func UserCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//首先获取cookie
		NameCookie, err := c.Cookie("user")
		if err != nil {
			log.Println("获取cookie失败 Error: ", err)
			c.Redirect(http.StatusMovedPermanently, "/login")
		}
		IsLoginCookie, err := c.Cookie("isLogin")
		if err != nil {
			log.Println("获取cookie失败 Error: ", err)
			c.Redirect(http.StatusMovedPermanently, "/login")
		} else {
			log.Println("获取cookie成功! IsLogin的值为： ", IsLoginCookie)
		}
		if IsLoginCookie != "true" {
			c.Redirect(http.StatusMovedPermanently, "/login")
		}

		//在redis中查找
		ans := util.CheckUserIsExists(NameCookie)
		if ans == false {
			log.Println("用户没有登录")
			c.Redirect(http.StatusMovedPermanently, "/login")
		}
		c.Next()

	}
}
