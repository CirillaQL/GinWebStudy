package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

//Upload 上传文件页面
func Upload(ctx *gin.Context) {
	ctx.HTML(200, "AddPhoto.html", gin.H{})
}

//GetPicture 获取
func GetPicture(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Fatal(err)
		return
	}
	username, _ := ctx.Cookie("user")
	ok := ctx.SaveUploadedFile(file, "./upload/"+username+"/"+file.Filename)
	if ok != nil {
		fmt.Println("保存时错误")
	} else {
		fmt.Println(file.Filename, "保存成功")
	}

}
