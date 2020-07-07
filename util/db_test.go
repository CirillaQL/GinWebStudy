package util

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"testing"
)

func TestLoginCheck(t *testing.T) {
	var name string = "Frankcox"
	db, err := sql.Open("mysql", "root:123456@/Web?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	//关闭数据库
	defer db.Close()

	var Phone string
	//查询
	var j string = "ql1194946223"
	result := db.QueryRow("select Password from User where Name = ?", name)
	err = result.Scan(&Phone)
	fmt.Println(result)
	fmt.Println(Phone)
	if err != nil {
		log.Println("登录失败，用户名不存在")
		t.Error(err)
	}

	res := bcrypt.CompareHashAndPassword([]byte(Phone), []byte(j))
	fmt.Println("验证密码结果: ", res)
}

func TestCheckUser(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@/Web?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	//关闭数据库
	defer db.Close()

	var user User
	var PhoneNumber string = "15840613358"
	result := db.QueryRow("select * from User where Phone = ?", PhoneNumber)
	err = result.Scan(&user.ID, &user.Name, &user.Password, &user.PhoneNumber)
	if err != nil {
		t.Error("err")
	}
	fmt.Println(user.ID)
	fmt.Println(user.Name)
	fmt.Println(user.Password)
	fmt.Println(user.PhoneNumber)
}

func TestReadPicture(t *testing.T) {
	srcDir := "../upload/Frankcox/"
	file, _ := ioutil.ReadDir(srcDir)
	for _, r := range file {
		fmt.Println(r.Name())
	}
}
