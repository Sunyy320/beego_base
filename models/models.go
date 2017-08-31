package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id int `orm:"auto"`
	Username string
	Pwd string
}

func (u *User) TableName() string  {
	return "user"
}

func init() {
	// 注册定义的model
	orm.RegisterModel(new (User))
}
