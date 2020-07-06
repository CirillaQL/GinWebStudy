package util

import (
	"log"
	"os"
	"testing"
)

func TestCreateFileDir(t *testing.T) {
	username := "test"
	DirPath := "../upload/" + username
	err := os.Mkdir(DirPath, os.ModePerm)
	if err != nil {
		log.Fatal("用户： ", username, "  创建文件夹失败， Error： ", err)
		//return false
	}
	//return true
}

func TestDeleteFileDir(t *testing.T) {
	username := "test"
	DirPath := "../upload/" + username
	err := os.RemoveAll(DirPath)
	if err != nil {
		log.Fatal("用户： ", username, "  删除文件夹失败, Error：  ", err)
	}
}
