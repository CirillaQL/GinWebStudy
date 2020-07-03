package controllers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

//Upload 上传文件页面
func Upload(ctx *gin.Context) {
	ctx.HTML(200, "upload.html", gin.H{})
}

//GetPicture 获取
func GetPicture(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	//fmt.Println(file.Size)
	if err != nil {
		log.Fatal(err)
		return
	}

	ok := ctx.SaveUploadedFile(file, "./upload/"+file.Filename)

	if ok != nil {
		fmt.Println("保存时错误")
	} else {
		fmt.Println(file.Filename, "保存成功")
	}

}
