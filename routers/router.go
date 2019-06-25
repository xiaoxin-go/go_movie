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
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/register", &controllers.RegisterController{})
    beego.Router("/moviecol", &controllers.MoviecolController{})
    beego.Router("/follow", &controllers.FollowController{})
    beego.Router("/logout", &controllers.LogoutController{})
}
