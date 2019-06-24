package controllers

import (
	"errors"
	"fmt"
	"goprj1/libs"
	"goprj1/models"
)

type PerformerController struct{
	libs.MyBeego
}

func (this *PerformerController) Get(){
	page, err := this.GetInt("page", 1)
	this.HttpParamsError(err, "page不能是字符")
	page_size, err := this.GetInt("page_size", 50)
	this.HttpParamsError(err, "page_size不能是字符")
	fmt.Printf("page:%s,page_size:%s,", page, page_size)

	var Performers []*models.Performer

	qs := O.QueryTable("performer")
	var count int64
	count, err = qs.Count()
	fmt.Println("count:",count)
	this.HttpServerError(err, "获取行数异常")
	_, err = qs.Limit(page*page_size, (page-1)*page_size).All(&Performers)
	this.HttpServerError(err, "获取数据异常")

	result_data :=  make(map[string]interface{})
	result_data["data_list"] = Performers
	result_data["total"] = count
	this.HttpSuccess(result_data, "数据请求成功")
}

func (this *PerformerController) Delete(){
	fmt.Println("this is delete")
	name := this.GetString("name")
	if name == ""{
		this.HttpParamsError(errors.New("name is empty"), "Params Error: name must not empty!")
	}
	Performer := models.Performer{Name: name}
	num, err := O.Delete(&Performer)
	qs := O.QueryTable("movie")
	num, err = qs.Filter("performer", name).Delete()
	fmt.Println(num)
	this.HttpServerError(err, "删除异常")
	fmt.Println(Performer)
	this.HttpSuccess(num, "删除成功")
}
