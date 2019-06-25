package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/orm"
	"goprj1/libs"
	"goprj1/models"
)

type FollowController struct{
	libs.MyBeego
}

func (this *FollowController)Get(){
	// 从session中获取用户ID
	user_id := this.GetSession("user_id")
	if user_id == nil{
		this.HttpServerError(errors.New("用户未登录"), "用户未登录")
	}
	// 获取请求page page_size
	this.GetPage()
	qs := O.QueryTable("follow")

	// 获取当前用户收藏的电影title
	var lists []orm.ParamsList
	_, err := qs.Filter("user_id", user_id).ValuesList(&lists, "performer")
	this.HttpServerError(err, "获取关注信息异常")

	// 更换表，获取电影数据，统计数量，写入到movies中
	qs = O.QueryTable("performer")
	qs = qs.Filter("name__in", lists)
	total, err := qs.Count()
	this.HttpServerError(err, "获取演员条数异常")

	var performers []models.Performer
	_, err = qs.Limit(this.Page*this.PageSize, (this.Page-1)*this.PageSize).All(&performers)
	this.HttpServerError(err, "关注演员写入数据异常")
	result_data := make(map[string]interface{})
	result_data["total"] = total
	result_data["performers"] = performers
	this.HttpSuccess(result_data, "数据请求成功")
}
type FollowParams struct{
	Performer string
}
func (this *FollowController)Post(){
	user_id := this.GetSession("user_id")
	if user_id == nil{
		this.HttpServerError(errors.New("用户未登录"), "用户未登录")
	}
	var params FollowParams
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	this.HttpParamsError(err, "获取参数异常")
	follow := models.Follow{Performer:params.Performer, UserId: user_id.(int)}
	id, err := O.Insert(&follow)
	this.HttpServerError(err, "添加关注失败")
	this.HttpSuccess(map[string]int{"id": int(id)}, "添加关注成功")
}
func(this *FollowController)Delete(){
	user_id := this.GetSession("user_id")
	if user_id == nil{
		this.HttpServerError(errors.New("用户未登录"), "用户未登录")
	}
	performer := this.GetString("performer")
	if performer == ""{
		this.HttpParamsError(errors.New("缺少必要参数"), "电影标题不能为空")
	}
	qs := O.QueryTable("follow")
	count, err := qs.Filter("performer", performer).Filter("user_id", user_id.(int)).Delete()
	this.HttpServerError(err, "删除关注异常")
	this.HttpSuccess(map[string]int{"total": int(count)}, "删除关注成功")
}
