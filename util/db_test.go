package util

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
)

func TestLoginCheck(t *testing.T) {
	var name string = "jo"
	var password string = "123456"
	db, err := sql.Open("mysql", "root:123456@/Web?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	//关闭数据库
	defer db.Close()

	var Password string
	//查询
	result := db.QueryRow("select * from User where Name = ?", name)
	err = result.Scan(&Password)
	fmt.Println(result)
	if err != nil {
		log.Println("登录失败，用户名不存在")
		t.Error(err)
	}
	if password == Password {
		t.Log("Ok")
	} else {
		t.Error(err)
	}
}
