package routers

import (
	"goprj1/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/mysql", &controllers.MysqlController{})
    beego.Router("/orm", &controllers.OrmController{})
    beego.Router("/test", &controllers.TestController{})
    beego.Router("/movie", &controllers.MovieController{})
}
