package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(context *gin.Context){
	context.HTML(http.StatusOK,"register.html",gin.H{
		"title": "Hello,begin to register",
	})
}
