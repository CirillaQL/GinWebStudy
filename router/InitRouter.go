package router

import (
	"GinWebStudy/controllers"
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
	router.StaticFile("/favicon.ico", "./static/icon/favicon.ico")

	//设置默认路由访问到错误网站时返回
	router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "Page Not Exists",
		})
	})

	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/upload", controllers.Upload)

	router.GET("/photo", func(context *gin.Context) {
		p := []string{"www.baidu.com", "www.baidu.com", "www.baidu.com"}
		context.HTML(http.StatusOK, "photo.html", gin.H{
			"name": "JoJo",
			"show": p,
		})
	})

	router.GET("/login", controllers.LoginGet)
	router.POST("/login", controllers.LoginPost)

	//注册
	router.GET("/register", controllers.RegisterGet)
	router.POST("/register", controllers.RegisterPost)

	work := router.Group("/api")
	{
		work.POST("/upload", controllers.GetPicture)
	}

	router.Run(":8080")
}
