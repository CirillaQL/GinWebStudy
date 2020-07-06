package util

import (
	uuid "github.com/satori/go.uuid"
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
