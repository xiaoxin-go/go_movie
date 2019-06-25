package controllers

import (
	"errors"
	"goprj1/libs"
)

type LogoutController struct{
	libs.MyBeego
}
func (this *LogoutController)Post(){
	id := this.GetSession("user_id")
	if id == nil{
		this.HttpServerError(errors.New("用户未登录"), "用户未登录")
	}
	this.DelSession("user_id")
	this.HttpSuccess(nil, "用户注销成功")
}