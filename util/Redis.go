package util

import (
	"fmt"
	_ "fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

//用于登录验证的Cookie结构体
type UserCookie struct {
	UserName string
	Password string
	LastTime int
}

//CreateUserCookie 创建用户Cookie
func CreateUserCookie(user User) UserCookie {
	cookie := UserCookie{
		UserName: user.Name,
		Password: user.Password,
		LastTime: 3 * 3600,
	}
	return cookie
}

//AddCookieToRedis 保存到Redis
func (cookie UserCookie) AddCookieToRedis() {
	//连接Redis
	con, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal("连接Redis 出错， Error: ", err)
	}
	defer con.Close()

	_, err = con.Do("hset", cookie.UserName, "password", cookie.Password)
	if err != nil {
		log.Fatal("用户: ", cookie.UserName, " 添加cookie失败！ Error: ", err)
	}
	_, err = con.Do("hset", cookie.UserName, "lasttime", cookie.LastTime)
	if err != nil {
		log.Fatal("用户: ", cookie.UserName, " 添加cookie失败！ Error: ", err)
	}
	_, err = con.Do("EXPIRE", cookie.UserName, cookie.LastTime)
	if err != nil {
		log.Fatal("用户: ", cookie.UserName, " 添加cookie失败！ Error: ", err)
	}
}

//CheckUserFromRedis 判断用户是否在Redis中
func CheckUserFromRedis(username string, password string) bool {
	//连接Redis
	con, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal("连接Redis 出错， Error: ", err)
	}
	defer con.Close()

	PasswordInRedis, err := redis.Bytes(con.Do("hget", username, "password"))
	if err != nil {
		log.Println("获取用户密码出错， Error: ", err)
		return false
	}

	fmt.Println(password)
	fmt.Println(string(PasswordInRedis))
	if password == string(PasswordInRedis) {
		return true
	} else {
		return false
	}
}

//SavePictureListInRedis 保存用户的图片信息到Redis中
