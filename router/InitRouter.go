package router

import (
	"GinWebStudy/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	//设置为ReleaseMode模式
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//加载文件
	router.LoadHTMLGlob("templates/*")
	router.Static("/static/css", "./static/css")
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

	router.POST("/", controllers.Login)

	router.GET("/register", controllers.RegisterGet)
	router.POST("/register", controllers.RegisterPost)

	work := router.Group("/api")
	{
		work.POST("/upload", controllers.GetPicture)
	}

	router.Run(":8080")
}
