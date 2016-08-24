package models

import (
	"errors"
	"log"
	"time"

	"Personal-Dictionary/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

// User Save User Info.
type User struct {
	Id            int64
	Username      string    `orm:"unique;size(32)" form:"Username"  valid:"Required;MaxSize(20);MinSize(6)"`
	Password      string    `orm:"size(64)" form:"Password" valid:"Required;MaxSize(20);MinSize(6)"`
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
	qs.Limit(pageSize, offset).OrderBy(sort).Values(&users, "Id", "Username", "Email", "Createtime", "Lastlogintime")
	count, _ = qs.Count()
	return users, count
}

// AddUser 添加用户
func AddUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}

	o := orm.NewOrm()
	user := new(User)
	user.Username = u.Username

	// Get New Salt and Handle Password.
	var saltLength int64
	if beego.AppConfig.String("salt_length") == "" {
		saltLength = 8
	} else {
		saltLength, _ = utils.Atoi64(beego.AppConfig.String("salt_length"))
	}
	userSalt := utils.GetNewSalt(saltLength)
	user.Salt = userSalt
	user.Password = utils.PassEncode(u.Password, user.Salt)
	user.Email = u.Email

	id, err := o.Insert(user)
	return id, err
}

// UpdateUser 更新用户信息
func UpdateUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}

	o := orm.NewOrm()
	user := make(orm.Params)
	if len(u.Username) > 0 {
		user["Username"] = u.Username
	}
	if len(u.Email) > 0 {
		user["Email"] = u.Email
	}

	var table User
	number, err := o.QueryTable(table).Filter("Id", u.Id).Update(user)
	return number, err
}

// DelUserById 删除用户
func DelUserById(Id int64) (int64, error) {
	o := orm.NewOrm()

	status, err := o.Delete(&User{Id: Id})
	return status, err
}

/********************* Danger
  Should Be Del in Produce Env.
   *************************/

// GetUserByUsername 获取用户信息（用于登录验证）
func GetUserByUsername(username string) (user User) {
	user = User{Username: username}
	o := orm.NewOrm()
	log.Println("user -->", user)

	o.Read(&user, "Username")
	return user
}
