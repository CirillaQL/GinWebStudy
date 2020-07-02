package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//IndexHTML 函数用于初始化主界面
func IndexHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
