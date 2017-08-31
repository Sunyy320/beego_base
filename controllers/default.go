package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego/orm"
	"mytest/models"
)

type MainController struct {
	beego.Controller
}

type User struct {
	Name string `form:"name" valid:"Required"`
	Age  int64  `form:"age"  valid:"Required;"`
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (this *MainController) Test() {
	// map [key_data_type] value_data_type
	queryString := this.GetString("name")
	logs.SetLogger("console")
	logs.Info("name=%s", queryString)

	// 设置session
	if queryString != "" {
		this.SetSession("name", queryString)
	}
	this.Data["json"] = map[string]interface{}{"code": 0, "msg": time.Now(), "name": this.GetSession("name")}
	this.ServeJSON()
	//this.Ctx.WriteString("pppp")
}

// 测试表单数据验证
func (this *MainController) ValidationTest() {
	u := User{}

	if err := this.ParseForm(&u); err != nil {
		this.Data["json"] = "error"
		this.ServeJSON()
		return
	}
	valid := validation.Validation{}
	// bool error
	b, err := valid.Valid(&u)
	if err != nil {
		this.Data["json"] = err
		this.ServeJSON()
		return
	}
	if !b {
		this.Data["json"] = valid.Errors
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{"code": 0, "msg": time.Now(), "data": "success"}
	this.ServeJSON()
}

// 测试插入数据库
func (this *MainController) DbInserTest() {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")

	user := new(models.User)
	user.Username = "hha"
	user.Pwd = "2233"

	id, err := o.Insert(user)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 3000, "msg": err}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{"code": 0, "msg": id, "res": "success"}
	this.ServeJSON()
}
