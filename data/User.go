package data

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
)

//用户结构体
type User struct {
	Id int
	Name string
	Password string
}

//生成用户id
func CreateID() int{
	rand.Seed(time.Now().Unix())
	return rand.Intn(10000)
}

//生成用户
func CreateUser(name string, password string) User{
	user := User{
		Id:       CreateID(),
		Name:     name,
		Password: password,
	}
	return user
}


