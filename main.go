package main

import (
	_ "mytest/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 注册驱动
	orm.RegisterDriver("mysql",orm.DRMySQL)
	// 设置默认数据库
	orm.RegisterDataBase("default","mysql","root:root@/mytest?charset=utf8")
}
func main() {
	beego.Run()
}

