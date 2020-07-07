package util

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
		return false
	}

	defer db.Close()

	//密码加密
	user.EncryptPassword()
	stmt, _ := db.Prepare("insert into Web.User(uuid,Name,Password,Phone) value (?,?,?,?);")
	_, err = stmt.Exec(user.ID, user.Name, user.Password, user.PhoneNumber)
	if err != nil {
		log.Println("插入数据库失败！ ERROR: ", err)
		return false
	}
	return true
}

//查询用户是否存在
func CheckUserExists(PhoneNumber string) bool {
	//第⼀步：打开数据库,格式是 ⽤户名：密码@/数据库名称？编码⽅式
	db, err := sql.Open("mysql", "root:123456@/Web?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	//关闭数据库
	defer db.Close()

	var user User
	result := db.QueryRow("select * from User where Phone = ?", PhoneNumber)
	err = result.Scan(&user.ID, &user.Name, &user.Password, &user.PhoneNumber)
	if err != nil {
		log.Println("查询失败！")
		return false
	}
	return true
}

//查询用户
func CheckUser(PhoneNumber string) User {
	//第⼀步：打开数据库,格式是 ⽤户名：密码@/数据库名称？编码⽅式
	db, err := sql.Open("mysql", "root:123456@/Web?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	//关闭数据库
	defer db.Close()

	var user User
	result := db.QueryRow("select * from User where Phone = ?", PhoneNumber)
	err = result.Scan(&user.ID, &user.Name, &user.Password, &user.PhoneNumber)
	if err != nil {
		log.Println("查询失败！")
		panic(err)
	}
	return user
}

//将本地文件存入数据库
func SavePictureListIntoDataBase(username string) {
	db, err := sql.Open("mysql", "root:123456@/Web?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	//关闭数据库
	defer db.Close()

	//stmt,err := db.Prepare("INSERT INTO WEB.Picture(NAME, Address, Owner, Describe) VALUES (?,?,?,?)")
	if err != nil {
		log.Fatal("Error: ", err)
	}

}
