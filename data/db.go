package data

import (
	"database/sql"
	"fmt"
	"log"
)

//将用户保存到数据库
func SaveToDataBase(user User) bool{
	//第⼀步：打开数据库,格式是 ⽤户名：密码@/数据库名称？编码⽅式
	db, err := sql.Open("mysql", "root:123456@/MyBlog?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	//关闭数据库
	defer db.Close()

	stmt, _ := db.Prepare("insert into User(ID,Name,Password) value (?,?,?);")
	_, err = stmt.Exec(user.Id,user.Name,user.Password)
	if err != nil{
		log.Println("插入数据库失败！ ERROR: ",err)
		return false
	}
	return true
}

//向数据库查询用户
func LoginCheck(name string, password string) bool{
	db, err := sql.Open("mysql", "root:123456@/MyBlog?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	//关闭数据库
	defer db.Close()

	var Password string
	//查询
	result := db.QueryRow("select Password from User where Name = ?",name)
	err = result.Scan(&Password);
	if err != nil{
		log.Println("登录失败，用户名不存在")
		return false
	}
	if password == Password{
		return true
	}else{
		return false
	}
}

