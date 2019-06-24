package controllers

import (
	"fmt"
	"goprj1/libs"
	"goprj1/models"
)

type MovieDetailController struct{
	libs.MyBeego
}

func (this *MovieDetailController) Get(){
	id, err := this.GetInt("id")
	if err != nil{
		this.HttpParamsError(err, "Params Error: id must is int")
	}
	fmt.Printf("id:%s \n", id)

	// 获取电影信息
	movie := models.Movie{Id: id}
	err = O.Read(&movie)
	this.HttpServerError(err, "获取电影信息异常")

	// 获取电影图片信息
	qs := O.QueryTable("image")
	var images []models.Image
	_, err =qs.Filter("title", movie.Title).All(&images)
	this.HttpServerError(err, "获取电影图片异常")

	// 获取电影链接信息
	qs = O.QueryTable("link")
	var links []models.Link
	_, err =qs.Filter("title", movie.Title).All(&links)
	this.HttpServerError(err, "获取电影图片异常")

	result_data :=  make(map[string]interface{})
	result_data["movie"] = movie
	result_data["images"] = images
	result_data["links"] = links
	this.HttpSuccess(result_data, "数据请求成功")
}
