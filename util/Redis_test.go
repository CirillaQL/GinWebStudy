package util

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"testing"
)

func TestCreateUserCookie(t *testing.T) {
	cookie := UserCookie{
		UserName: "Frankcox",
		Password: "ql1194946223",
		LastTime: 3600,
	}
	fmt.Println(cookie)
}

func TestUserCookie_AddCookieToRedis(t *testing.T) {
	cookie := UserCookie{
		UserName: "Frankcox",
		Password: "ql1194946223",
		LastTime: 600,
	}
	cookie.AddCookieToRedis()
}

func TestCheckCookie(t *testing.T) {
	con, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal("连接Redis 出错， Error: ", err)
	}
	defer con.Close()

	ans, err := con.Do("EXISTS", "Frankcox")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ans)
}
