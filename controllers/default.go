package controllers

import (
	"github.com/astaxie/beego"
)

type user struct {
	Name string `form:"username"`
	Age int
	Sex string
}
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	//自定义开发
	c.Data["myName"]="xianmeng.jiang"
	u :=&user{"xm.j",30,"男"}
	c.Data["user"] =u
}

//Post 方式
func (this *MainController)Post(){

	u := user{}
	if err := this.ParseForm(&u); err !=nil{
		beego.Info(u.Name,u.Age,u.Sex)
	}
	beego.Info("正常获取",u.Name,u.Age,u.Sex)
	this.Data["user"]= &u

	this.TplName ="formPost.tpl"
	this.Data["html"]="<div> this is html tag </div>"
}
