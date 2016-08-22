package models

import (
	"time"

	"github.com/astaxie/beego"
)

type User struct {
	Id         int64
	Username   string
	Password   string
	Salt       string
	Createtime time.Time
}

// TableName User Table Name.
func (u *User) TableName() string {
	return beego.AppConfig.String("db_user_table")
}
