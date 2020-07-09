package util

import (
	"fmt"
	"testing"
)

func TestGetPictureList(t *testing.T) {
	var username string = "Frankcox"
	s := GetPictureListFromDir(username)
	fmt.Println(s)
}

func TestSavePictureList(t *testing.T) {
	var username string = "Frankcox"
	s := GetPictureListFromDir(username)
	fmt.Println(s)
	SavePictureListIntoDataBase(s)
}
