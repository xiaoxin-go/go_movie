package controllers

import (
	"fmt"
	"goprj1/libs"
	"goprj1/models"
)

type PerformerDetailController struct{
	libs.MyBeego
}

func (this *PerformerDetailController) Get(){
	id, err := this.GetInt("id")
	this.HttpParamsError(err, "id is must int")
	page, err := this.GetInt("page", 1)
	this.HttpParamsError(err, "page不能是字符")
	page_size, err := this.GetInt("page_size", 50)
	this.HttpParamsError(err, "page_size不能是字符")
	fmt.Printf("id:%s \n", id, page, page_size)

	// 获取演员信息
	performer := models.Performer{Id: id}
	err = O.Read(&performer)
	this.HttpServerError(err, "获取演员信息异常")

	// 获取演员电影信息
	qs := O.QueryTable("movie")
	var movies []models.Movie
	qs = qs.Filter("performer__contains", performer.Name)
	total, err := qs.Count()
	this.HttpServerError(err, "获取电影条数异常")
	_, err = qs.Limit(page*page_size, (page-1)*page_size).All(&movies)
	this.HttpServerError(err, "获取演员电影异常")

	result_data :=  make(map[string]interface{})
	result_data["performer"] = performer
	result_data["movies"] = movies
	result_data["total"] = total
	this.HttpSuccess(result_data, "数据请求成功")
}