package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type OrmController struct{
	beego.Controller
}

type User struct{
	Id int
	Name string
}

func (self *OrmController) Get(){
	// 1. 连接数据库
	orm.RegisterDataBase("default", "mysql", "root:xiaoxin@tcp(127.0.0.1:3306)/video?charset=utf8")

	// 2. 注册表
	orm.RegisterModel(new(User))

	// 3.生成表
	orm.RunSyncdb("default", false, true)
	self.Ctx.WriteString("创建表成功")

	o := orm.NewOrm()
	user := User{}
	user.Name = "hanmeimei"
	_, err := o.Insert(&user)
	if err != nil{
		beego.Info("插入失败： ", err)
		return
	}
	self.Ctx.WriteString("插入成功")

	o1 := orm.NewOrm()

	user1 := User{Id:1}

	err = o1.Read(&user1)
	if err != nil{
		beego.Info("查询失败： ", err)
		return
	}

	user.Name = "lilei"
	_, err = o1.Update(&user)
	if err != nil{
		beego.Info("更新数据错误")
		return
	}
	self.Ctx.WriteString("更新成功")

}