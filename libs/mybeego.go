package libs

import "github.com/astaxie/beego"

type MyBeego struct{
	beego.Controller
	HttpResponse
}
func (this *MyBeego)HttpSuccess(result_data interface{}, message string){
	this.Data["json"] = this.Success(result_data, "数据请求成功")
	this.ServeJSON()
}
func (this *MyBeego)HttpServerError(err error, message string){
	if err != nil{
		beego.Info(err)
		this.Data["json"] = this.ServerError(message)
		this.ServeJSON()
	}
}
func (this *MyBeego)HttpParamsError(err error, message string){
	if err != nil{
		beego.Info(err)
		this.Data["json"] = this.ParamsError(message)
		this.ServeJSON()
	}
}