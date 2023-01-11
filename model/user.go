package model

import (
	"github.com/archine/ioc"
	log "github.com/sirupsen/logrus"
)

// User 和数据库对应得结构体
type User struct {
	Age  int
	Name string
}

// UserMapper 操作数据库层
type UserMapper struct{}

func (u *UserMapper) CreateBean() ioc.Bean {
	return &UserMapper{}
}

// AddUser 添加用户
func (u *UserMapper) AddUser(arg *User) {
	log.Printf("%v\n", arg)
}

// GetUser 获取用户
func (u *UserMapper) GetUser() *User {
	return &User{10, "张三"}
}
