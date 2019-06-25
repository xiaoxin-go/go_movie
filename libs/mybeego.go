package libs

import (
	"github.com/astaxie/beego"
)

type MyBeego struct{
	beego.Controller
	HttpResponse
	Page int
	PageSize int
}
func (this *MyBeego)HttpSuccess(result_data interface{}, message string){
	this.Data["json"] = this.Success(result_data, message)
	this.ServeJSON()
	this.StopRun()
}
func (this *MyBeego)HttpServerError(err error, message string){
	if err != nil{
		beego.Info(err)
		this.Data["json"] = this.ServerError(message)
		this.ServeJSON()
		this.StopRun()
	}
}
func (this *MyBeego)HttpParamsError(err error, message string){
	if err != nil{
		beego.Info(err)
		this.Data["json"] = this.ParamsError(message)
		this.ServeJSON()
		this.StopRun()
	}
}
func (this *MyBeego)GetPage(){
	page, err := this.GetInt("page", 1)
	this.HttpParamsError(err, "page不能是字符")
	page_size, err := this.GetInt("page_size", 50)
	this.HttpParamsError(err, "page_size不能是字符")
	this.Page = page
	this.PageSize = page_size
}