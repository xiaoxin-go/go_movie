package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"goprj1/libs"
	"goprj1/models"
)

type MovieController struct{
	libs.MyBeego
}

func (this *MovieController) Get(){
	keyword := this.GetString("keyword")
	performer := this.GetString("performer")
	genre := this.GetString("genre")
	series := this.GetString("series")
	page, err := this.GetInt("page", 1)
	this.HttpParamsError(err, "page不能是字符")
	page_size, err := this.GetInt("page_size", 50)
	this.HttpParamsError(err, "page_size不能是字符")
	fmt.Printf("keyword:%s,performer:%s,genre:%s,series:%s,page:%s,page_size:%s,",
		keyword, performer, genre, series, page, page_size)

	var movies []*models.Movie

	qs := O.QueryTable("movie")
	var count int64
	switch{
	case keyword != "":
		cond := orm.NewCondition()
		cond1 := cond.And("title__contains", keyword).Or("info__contains", keyword)
		qs = qs.SetCond(cond1)
	case performer != "":
		qs = qs.Filter("performer__contains", performer)
	case series != "":
		qs = qs.Filter("series", series)
	case genre != "":
		qs = qs.Filter("genre", genre)
	}
	count, err = qs.Count()
	fmt.Println(count)
	this.HttpServerError(err, "获取行数异常")
	_, err = qs.Limit(page*page_size, (page-1)*page_size).All(&movies)
	this.HttpServerError(err, "获取数据异常")

	result_data :=  make(map[string]interface{})
	result_data["data_list"] = movies
	result_data["total"] = count
	this.HttpSuccess(result_data, "数据请求成功")
}

func (this *MovieController) Delete(){
	fmt.Println("this is delete")
	id, err := this.GetInt("id")
	this.HttpParamsError(err, "Params Error: Id get error!")
	movie := models.Movie{Id: id}
	num, err := O.Delete(&movie)
	qs := O.QueryTable("movie")
	num, err = qs.Filter("id__gt", 7).Delete()
	fmt.Println(num)
	this.HttpServerError(err, "删除异常")
	fmt.Println(movie)
	this.HttpSuccess(num, "删除成功")
}