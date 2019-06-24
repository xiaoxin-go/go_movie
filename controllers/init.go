package controllers

import (
	"github.com/astaxie/beego/orm"
	"goprj1/models"
)

var O orm.Ormer

func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:xiaoxin@tcp(127.0.0.1:3306)/movie?charset=utf8", 30)
	orm.RegisterModel(new(models.Movie))
	orm.RegisterModel(new(models.Performer))
	orm.RegisterModel(new(models.Link))
	orm.RegisterModel(new(models.Image))
	O = orm.NewOrm()
	O.Using("default")
}