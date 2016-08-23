package models

import (
	"errors"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

// User Save User Info.
type User struct {
	Id            int64
	Username      string    `orm:"unique;size(32)" form:"Username"  valid:"Required;MaxSize(20);MinSize(6)"`
	Password      string    `orm:"size(32)" form:"Password" valid:"Required;MaxSize(20);MinSize(6)"`
	Repassword    string    `orm:"-" form:"Repassword" valid:"Required"`
	Salt          string    `orm:"size(32)" form:"Salt" valid:"Required;MaxSize(20);MinSize(6)"`
	Email         string    `orm:"size(32)" form:"Email" valid:"Email"`
	Lastlogintime time.Time `orm:"null;type(datetime)" form:"-"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add" `
}

// TableName User Table Name.
func (u *User) TableName() string {
	return beego.AppConfig.String("db_user_table")
}

func init() {
	orm.RegisterModel(new(User))
}

// Valid 密码验证
func (u *User) Valid(v *validation.Validation) {
	if u.Password != u.Repassword {
		v.SetError("Repassword", "两次输入的密码不一样")
	}
}

// checkUser 检查用户信息合法性
func checkUser(u *User) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&u)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

/************************************************************/

// GetUserList 返回所有用户信息列表
func GetUserList(page int64, pageSize int64, sort string) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * pageSize
	}
	qs.Limit(pageSize, offset).OrderBy(sort).Values(&users)
	count, _ = qs.Count()
	return users, count
}
