package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"goprj1/libs"
	"goprj1/models"
)

type RegisterController struct{
	libs.MyBeego
}

type RegisterParams struct{
	Username string
	Password string
	Nickname string
}

func (this *RegisterController) Post(){
	var params RegisterParams
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	this.HttpParamsError(err, "获取参数异常")
	if params.Username == "" || params.Password == "" || params.Nickname == ""{
		this.HttpServerError(errors.New("用户名或密码或昵称不能为空"), "用户名或密码或昵称不能为空")
	}
	qs := O.QueryTable("user")
	user_exits := qs.Filter("username", params.Username).Exist()
	if user_exits{
		this.HttpServerError(errors.New("用户名已存在"), "用户名已存在")
	}
	var user models.User
	user.Username = params.Username
	user.Password = params.Password
	user.Nickname = params.Nickname
	id, err :=O.Insert(&user)
	fmt.Println(id)
	this.HttpServerError(err, "注册用户失败")
	this.HttpSuccess(map[string]int{"Id":int(id)}, "用户注册成功")
}