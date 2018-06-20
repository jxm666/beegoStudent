package main

import (
	_ "studentLog/routers"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"fmt"
	"studentLog/models"
)

/*func init() {
	// 注册数据库
	models.RegisterDB()
}*/

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//这块有点坑 需要tcp()起来些ip与端口
	orm.RegisterDataBase("default", "mysql", "root:ZXdc.yjsb12@tcp(192.168.1.157:3306)/orm_db?charset=utf8")
}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

/*	profile := new(models.Profile)
	profile.Age = 30

	user := new(models.User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))*/


	//CRUD 操作演示
	user := &models.MyUser{Name:"test 3"}
	//new 对象返回是对象的指针地址
	/*user := new (models.MyUser)
	user.Name="test "*/

	//fmt.Println(user)
	fmt.Println(o.Insert(user))
	user.Name = "updat test3"

	fmt.Println(o.Update(user))
	fmt.Println(o.Read(user))
	fmt.Println(o.Delete(user))

	beego.Run()
}
