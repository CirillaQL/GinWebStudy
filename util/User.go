package util

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
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
	var OriginalPassword string = user.Password
	EncryptionPassword, err := bcrypt.GenerateFromPassword([]byte(OriginalPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("用户: ", user.Name, " 加密失败! Error: ", OriginalPassword)
		return
	}
	AfterPassword := string(EncryptionPassword)
	user.Password = AfterPassword
}

/*
	CheckPassword
	匹配用户密码，将输入密码与数据库中密码进行匹配
*/
func (user *User) CheckPassword(InputPassword string) bool {
	result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(InputPassword))
	if result != nil {
		return false
	}
	return true
}
