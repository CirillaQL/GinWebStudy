package router

import (
	"GinWebStudy/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Photo struct {
	Name string
	Url  string
}

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

	//路由分组
	HomePage := router.Group("/homepage")
	{
		HomePage.GET("/:suer")
	}

	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/upload", controllers.Upload)

	router.GET("/photo", func(context *gin.Context) {
		a := Photo{
			Name: "1",
			Url:  "../static/img/2k5zd9.jpg",
		}
		b := Photo{
			Name: "2",
			Url:  "../static/img/2k5zd9.jpg",
		}

		k := []Photo{a, b}
		//p := []string{"../static/img/2k5zd9.jpg", "../static/img/2k5zd9.jpg", "../static/img/2k5zd9.jpg"}
		context.HTML(http.StatusOK, "photo.html", gin.H{
			"name":  "你好,JoJo",
			"image": k,
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
