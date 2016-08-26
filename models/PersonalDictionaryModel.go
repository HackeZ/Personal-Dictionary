package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// PersonalDictionary Save Personal Dictionary Info.
type PersonalDictionary struct {
	Id      int64
	User    *User  `orm:"rel(fk);size(32)" form:"User"  valid:"Required;MaxSize(20);MinSize(6)"`
	Keyword string `orm:"size(256);index" form:"Keyword" valid:"Required;MaxSize(256);MinSize(0)"`
	Content string `orm:"size(256)" form:"Content" valid:"MaxSize(256);MinSize(0)"`
	// Tags string
	Createtime time.Time `orm:"type(datetime);auto_now_add"`
}

// TableName Personal-Dictionary Table Name.
func (pd PersonalDictionary) TableName() string {
	return beego.AppConfig.String("db_personaldictionary_table")
}

func init() {
	orm.RegisterModel(new(PersonalDictionary))
}

// checkPD 确保每个用户只有一个相同 Keyword 的数据
func checkPD(pd *PersonalDictionary) (err error) {
	o := orm.NewOrm()
	var oldPD PersonalDictionary
	err = o.QueryTable(pd).Filter("Keyword", pd.Keyword).One(&oldPD)
	// 没有找到纪录
	if err == orm.ErrNoRows {
		return nil
	}
	err = errors.New("你已经有这条词典了～ 找到并更新它吧！")
	return
}

/************************************************************/

// GetPersonalDictionaryList 获取个人字典列表
func GetPersonalDictionaryList(page int64, pageSize int64, sort string) (pds []orm.Params, count int64) {
	o := orm.NewOrm()
	pd := new(PersonalDictionary)
	qs := o.QueryTable(pd)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * pageSize
	}
	qs.Limit(pageSize, offset).OrderBy(sort).Values(&pds, "Keyword", "Content", "Createtime")
	count, _ = qs.Count()
	return pds, count
}

// AddPersonalDictionary 添加一条个人词典
func AddPersonalDictionary(pd *PersonalDictionary) (int64, error) {
	if err := checkPD(pd); err != nil {
		return 0, err
	}

	o := orm.NewOrm()
	perdic := new(PersonalDictionary)

	perdic.User = pd.User
	perdic.Keyword = pd.Keyword
	perdic.Content = pd.Content

	id, err := o.Insert(perdic)
	return id, err
}

// UpdatePersonalDictionary 更新个人词典数据
func UpdatePersonalDictionary(pd *PersonalDictionary) (int64, error) {
	if err := checkPD(pd); err != nil {
		return 0, err
	}

	o := orm.NewOrm()
	perdic := make(orm.Params)

	if len(pd.Keyword) > 0 {
		perdic["Keyword"] = pd.Keyword
	}
	if len(pd.Content) > 0 {
		perdic["Content"] = pd.Content
	}

	var table PersonalDictionary
	number, err := o.QueryTable(table).Filter("Id", pd.Id).Update(perdic)
	return number, err
}

// DelPersonalDictionary 根据ID删除一条个人词典数据
func DelPersonalDictionary(Id int64) (int64, error) {
	o := orm.NewOrm()

	status, err := o.Delete(&PersonalDictionary{Id: Id})
	return status, err
}
