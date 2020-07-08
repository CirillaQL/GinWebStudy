package util

import (
	"log"
	"os"
)

//CreateFileDir 给用户创建文件夹
func CreateFileDir(username string) bool {
	DirPath := "./upload/" + username
	err := os.Mkdir(DirPath, os.ModePerm)
	if err != nil {
		log.Fatal("用户： ", username, "  创建文件夹失败， Error： ", err)
		return false
	}
	return true
}

//DeleteFileDir 删除文件夹
func DeleteFileDir(username string) bool {
	DirPath := "./upload/" + username
	err := os.RemoveAll(DirPath)
	if err != nil {
		log.Fatal("用户： ", username, "  删除文件夹失败, Error：  ", err)
		return false
	}
	return true
}
