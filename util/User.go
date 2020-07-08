package util

import (
	"encoding/base64"
	uuid "github.com/satori/go.uuid"
	"log"
)

//User 用户结构体
type User struct {
	ID          string
	Name        string
	PhoneNumber string
	Password    string
}

//CreateID 生成uuid
func CreateID() string {
	u1 := uuid.NewV4()
	uuid := u1.String()
	return uuid
}

//CreateUser 生成用户
func CreateUser(name string, password string, number string) User {
	user := User{
		ID:          CreateID(),
		Name:        name,
		Password:    password,
		PhoneNumber: number,
	}
	return user
}

/*
	EncryptPassword
	将用户的密码加密，保存到数据库中
*/
func (user *User) EncryptPassword() {
	input := []byte(user.Password)
	encodeString := base64.StdEncoding.EncodeToString(input)
	user.Password = encodeString
}

/*
	CheckPassword
	匹配用户密码，将输入密码与数据库中密码进行匹配
*/
func (user *User) CheckPassword(InputPassword string) bool {
	decodeBytes, err := base64.StdEncoding.DecodeString(user.Password)
	if err != nil {
		log.Fatalln(err)
	}
	aaa, err := base64.StdEncoding.DecodeString(string(decodeBytes))
	if err != nil {
		log.Fatalln(err)
	}
	if InputPassword == string(aaa) {
		return true
	} else {
		return false
	}
}
