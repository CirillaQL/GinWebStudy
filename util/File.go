package util

import (
	"io/ioutil"
	"log"
	"os"
)

//CreateFileDir 给用户创建文件夹
func CreateFileDir(username string) bool {
	DirPath := "./upload/" + username
	os.Mkdir(DirPath, os.ModePerm)
	//if err != nil {
	//log.Fatal("用户： ", username, "  创建文件夹失败， Error： ", err)
	//return false
	//}
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

//GetPictureList 从文件夹获取用户的图片
func GetPictureListFromDir(username string) []Picture {
	var PictureList []Picture
	srcDir := "../upload/" + username + "/"
	file, _ := ioutil.ReadDir(srcDir)
	for _, r := range file {
		Path := srcDir[1:] + r.Name()
		picture := Picture{
			Name:     r.Name(),
			Describe: "描述",
			Address:  Path,
			Owner:    username,
		}
		PictureList = append(PictureList, picture)
	}
	return PictureList
}
