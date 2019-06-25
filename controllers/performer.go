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
	this.GetPage()

	var Performers []*models.Performer

	qs := O.QueryTable("performer")
	var count int64
	count, err := qs.Count()
	fmt.Println("count:",count)
	this.HttpServerError(err, "获取行数异常")
	_, err = qs.Limit(this.Page*this.PageSize, (this.Page-1)*this.PageSize).All(&Performers)
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
