package router

import (
	"GinWebStudy/controllers"
	"GinWebStudy/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	//设置为ReleaseMode模式
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//加载文件
	router.LoadHTMLGlob("templates/*")
	router.Static("/static/css", "./static/css")
	router.Static("/static/img", "./static/img")
	router.Static("/static/libs", "./static/libs")
	router.Static("/static/js", "./static/js")
	router.Static("/homepage/upload", "./upload")
	router.StaticFile("/favicon.ico", "./static/icon/favicon.ico")

	//登录
	router.GET("/login", controllers.LoginGet)
	router.POST("/login", controllers.LoginPost)
	router.GET("/loginWrong", controllers.LoginWrong)

	//注册
	router.GET("/register", controllers.RegisterGet)
	router.POST("/register", controllers.RegisterPost)

	//路由分组
	HomePage := router.Group("/homepage")
	HomePage.Use(middleware.UserCheck())
	{
		HomePage.GET("/", controllers.IndexHTML)
	}

	router.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/login")
	})

	router.GET("/upload", controllers.Upload)

	work := router.Group("/api")
	{
		work.POST("/upload", controllers.GetPicture)
	}

	router.Run(":8080")
}
