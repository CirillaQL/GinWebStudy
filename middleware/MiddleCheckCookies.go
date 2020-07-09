package middleware

import (
	"GinWebStudy/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//登录检查中间件
func UserCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//首先获取cookie
		GetNameFromCookie, err := c.Cookie("user")
		if err != nil {
			log.Println("获取cookie失败 Error: ", err)
			c.Redirect(http.StatusMovedPermanently, "/login")
		}
		GetPasswordFromCookie, err := c.Cookie("password")
		if err != nil {
			log.Println("获取cookie失败 Error: ", err)
			c.Redirect(http.StatusMovedPermanently, "/login")
		}
		//在redis中查找
		ans := util.CheckUserFromRedis(GetNameFromCookie, GetPasswordFromCookie)
		fmt.Println(ans)
		if ans {
			c.Next()
		} else {
			c.Redirect(http.StatusMovedPermanently, "/login")
		}
	}
}
