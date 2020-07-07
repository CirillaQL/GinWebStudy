package util

import (
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

//CheckUserIsExists 判断用户是否在Redis中
func CheckUserIsExists(username string) bool {
	//连接Redis
	con, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal("连接Redis 出错， Error: ", err)
	}
	defer con.Close()

	ans, err := con.Do("EXISTS", username)
	if err != nil {
		log.Fatal("在查询用户是否在Redis时出错，Error: ", err)
	}
	if ans == 0 {
		return false
	} else {
		return true
	}
}
