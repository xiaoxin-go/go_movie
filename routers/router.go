package routers

import (
	"goprj1/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/movie", &controllers.MovieController{})
    beego.Router("/performer", &controllers.PerformerController{})
    beego.Router("/movie/detail", &controllers.MovieDetailController{})
    beego.Router("/performer/detail", &controllers.PerformerDetailController{})
}
