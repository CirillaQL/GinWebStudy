package controllers

import (
	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	ctx.HTML(200, "upload.html", gin.H{})
}
