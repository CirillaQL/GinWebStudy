package util

import (
	"fmt"
	"testing"
)

func TestGetPictureList(t *testing.T) {
	var username string = "Frankcox"
	s := GetPictureList(username)
	fmt.Println(s)
}
