package util

import (
	"encoding/base64"
	"fmt"
	"log"
	"testing"
)

func TestUser(t *testing.T) {
	s := CreateID()
	fmt.Println(s)
}

func TestUser_CheckPassword(t *testing.T) {
	a := User{
		ID:          "2321321",
		Name:        "Frankcox",
		PhoneNumber: "15840613358",
		Password:    "ql1194946223",
	}
	input := []byte(a.Password)
	encodeString := base64.StdEncoding.EncodeToString(input)

	fmt.Println("加密后的结果: ", encodeString)

	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("解密后: ", string(decodeBytes))

}
