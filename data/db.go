package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//SaveToDataBase 将用户保存到数据库
func SaveToDataBase(user User) bool {
	//第⼀步：打开数据库,格式是 ⽤户名：密码@/数据库名称？编码⽅式
	db, err := sql.Open("mysql", "root:123456@/Web?charset=utf8")
	if err != nil {
		log.Fatal("数据库打开错误!  Error: ", err)
	}
	//关闭数据库
	defer db.Close()

	stmt, _ := db.Prepare("insert into User(uuid,Name,Password,Phone) value (?,?,?,?);")
	_, err = stmt.Exec(user.ID, user.Name, user.Password, user.PhoneNumber)
	if err != nil {
		log.Println("插入数据库失败！ ERROR: ", err)
		return false
	}
	return true
}

//查询用户名是否存在
func CheckUser(username string) bool {
	//第⼀步：打开数据库,格式是 ⽤户名：密码@/数据库名称？编码⽅式
	db, err := sql.Open("mysql", "root:123456@/Web?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	//关闭数据库
	defer db.Close()

	//查询
	var Password string
	result := db.QueryRow("select password from User where Name = ?", username)
	err = result.Scan(&Password)
	if err == nil {
		log.Println("用户名重复，注册失败")
		return false
	}
	return true
}
