package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goprj1/models"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
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
	page := this.GetString("page")
	page_size := this.GetString("page_size")

	page_int, err := strconv.Atoi(page)
	if err != nil{
		fmt.Println("TypeError: page not is int: ", page, page_int)
		this.Ctx.WriteString("TypeError: page not is int")
		return
	}

	page_size_int, err := strconv.Atoi(page_size)
	if err != nil{
		fmt.Println("TypeError: page_size not is int: ", page_size, page_size_int)
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
		_, err = qs.Limit(page, page_int*page_size_int).All(&movies)
		if err != nil{
			fmt.Println("获取数据异常：", err, count)
			this.Ctx.WriteString("获取数据异常：")
			return
		}
	}

	this.Ctx.WriteString("this is get")
}

func (this *MovieController) Delete(){
	fmt.Println("this is delete")
	this.Ctx.WriteString("this is delete")
}