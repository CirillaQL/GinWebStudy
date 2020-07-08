package util

import (
	"io/ioutil"
	"log"
	"os"
)

//Picture 结构体
type Picture struct {
	Name     string
	Describe string
	Address  string
	Owner    string
}

//ChangeName 修改文件名
func (p *Picture) ChangeName(username string, filename string) bool {
	file := "./upload/" + username + "/" + p.Name
	err := os.Rename(file, "../upload/"+username+"/"+filename)
	if err != nil {
		log.Fatal("用户: ", username, " 修改图片名： ", p.Name, " 失败")
		return false
	}
	return true
}

func GetPictureList(username string) []Picture {
	var PictureList []Picture
	srcDir := "./upload/" + username + "/"
	file, _ := ioutil.ReadDir(srcDir)
	for _, r := range file {
		Path := srcDir + r.Name()
		picture := Picture{
			Name:     r.Name(),
			Describe: "",
			Address:  Path,
			Owner:    username,
		}
		PictureList = append(PictureList, picture)
	}
	return PictureList
}
