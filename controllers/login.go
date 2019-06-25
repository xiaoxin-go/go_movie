package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"goprj1/libs"
	"goprj1/models"
)

type LoginController struct{
	libs.MyBeego
}

type LoginParams struct{
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (this *LoginController) Post(){
	var params LoginParams
	var user models.User
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	this.HttpServerError(err, "获取参数异常")
	fmt.Println(params)
	if params.UserName == "" || params.Password == ""{
		this.HttpServerError(errors.New("用户名或密码不能为空"), "用户名或密码不能为空")
	}
	qs := O.QueryTable("user")
	err = qs.Filter("username", params.UserName).One(&user)
	this.HttpServerError(err, "用户信息获取异常")
	fmt.Println(user)
	if user.Password != params.Password{
		this.HttpServerError(errors.New("用户名密码不正确"), "用户名或密码不正确")
	}
	this.SetSession("user_id", user.Id)
	this.HttpSuccess(params, "登录成功")
}
