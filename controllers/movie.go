package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goprj1/models"
	_ "github.com/go-sql-driver/mysql"
)

func init(){
	orm.RegisterDataBase("default", "mysql", "root:xiaoxin@tcp(127.0.0.1:3306)/movie?charset=utf8", 30)
}

type MovieController struct{
	beego.Controller
}

func (this *MovieController) Get(){
	keyword := this.GetString("keyword")
	performer := this.GetString("performer")
	genre := this.GetString("genre")
	series := this.GetString("series")
	page, err := this.GetInt("page", 1)
	if err != nil{
		fmt.Println("TypeError: page not is int: ", page)
		this.Ctx.WriteString("TypeError: page not is int")
		return
	}
	page_size, err := this.GetInt("page_size", 50)
	if err != nil{
		fmt.Println("TypeError: page_size not is int: ", page_size)
		this.Ctx.WriteString("TypeError: page not is int")
		return
	}

	fmt.Println(keyword, performer, genre, series, page, page_size)
	orm.RegisterDataBase("default", "mysql", "root:xiaoxin@tcp(127.0.0.1:3306)/movie?charset=utf8", 30)

	var movies []*models.Movie
	o := orm.NewOrm()
	qs := o.QueryTable("movie")
	if keyword != ""{
		cond := orm.NewCondition()
		cond1 := cond.And("title__contains", keyword).Or("info__contains", keyword)
		qs := qs.SetCond(cond1)
		count, err := qs.Count()
		if err != nil{
			fmt.Println("获取行数异常：", err, count)
			this.Ctx.WriteString("获取行数异常：")
			return
		}
		_, err = qs.Limit(page, page*page_size).All(&movies)
		if err != nil{
			fmt.Println("获取数据异常：", err, count)
			this.Ctx.WriteString("获取数据异常：")
			return
		}
	}
	this.Data["json"] = &movies
	this.ServeJSON()
}

func (this *MovieController) Delete(){
	fmt.Println("this is delete")
	id, err := this.GetInt("id")
	if err != nil{
		beego.Info("Params Error: Id get error!", err, id)
		this.ServeJSON()
	}
	
	mystruct := models.User{Id:id}
	this.Data["json"] = &mystruct
	this.ServeJSON()
}