package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"goprj1/libs"
	"goprj1/models"
)

type MoviecolController struct{
	libs.MyBeego
}

func (this *MoviecolController)Get(){
	// 从session中获取用户ID
	user_id := this.GetSession("user_id")
	if user_id == nil{
		this.HttpServerError(errors.New("用户未登录"), "用户未登录")
	}
	// 获取请求page page_size
	this.GetPage()
	qs := O.QueryTable("moviecol")

	// 获取当前用户收藏的电影title
	var lists []orm.ParamsList
	_, err := qs.Filter("user_id", user_id).ValuesList(&lists, "title")
	this.HttpServerError(err, "获取收藏信息异常")

	// 更换表，获取电影数据，统计数量，写入到movies中
	qs = O.QueryTable("movie")
	qs = qs.Filter("title__in", lists)
	total, err := qs.Count()
	this.HttpServerError(err, "获取电影条数异常")

	var movies []models.Movie
	_, err = qs.Limit(this.Page*this.PageSize, (this.Page-1)*this.PageSize).All(&movies)
	this.HttpServerError(err, "收藏电影写入数据异常")
	result_data := make(map[string]interface{})
	result_data["total"] = total
	result_data["movies"] = movies
	this.HttpSuccess(result_data, "数据请求成功")
}
type MovieColParams struct{
	Title string
}
func (this *MoviecolController)Post(){
	user_id := this.GetSession("user_id")
	if user_id == nil{
		this.HttpServerError(errors.New("用户未登录"), "用户未登录")
	}
	fmt.Println(".....................")
	var params MovieColParams
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &params)
	fmt.Println("err:",err, params)
	this.HttpParamsError(err, "获取参数异常")
	if params.Title == ""{
		this.HttpParamsError(errors.New("title不能为空"), "title不能为空")
	}
	moviecol := models.MovieCol{Title:params.Title, UserId: user_id.(int)}
	id, err := O.Insert(&moviecol)
	this.HttpServerError(err, "添加收藏失败")
	this.HttpSuccess(map[string]int{"id": int(id)}, "添加收藏成功")
}
func(this *MoviecolController)Delete(){
	user_id := this.GetSession("user_id")
	if user_id == nil{
		this.HttpServerError(errors.New("用户未登录"), "用户未登录")
	}
	title := this.GetString("title")
	if title == ""{
		this.HttpParamsError(errors.New("缺少必要参数"), "电影标题不能为空")
	}
	qs := O.QueryTable("moviecol")
	count, err := qs.Filter("title", title).Filter("user_id", user_id.(int)).Delete()
	this.HttpServerError(err, "删除收藏异常")
	this.HttpSuccess(struct{Total int `json:"total"`}{Total: int(count)}, "删除收藏成功")
}
