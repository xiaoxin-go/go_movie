package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type TestController struct{
	beego.Controller
}

func (self *TestController) Get(){
	beego.Info("this is get")
	self.Ctx.WriteString("this is get")
}

func (self *TestController) Post(){
	name := self.GetString("name")
	fmt.Println("name is :", name)
	beego.Info("this is post")
	self.Ctx.WriteString("this is post")
}