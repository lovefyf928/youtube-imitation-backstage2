package routers

import (
	"youtube-imitation-backstage2/controllers"
	"github.com/astaxie/beego"
)



func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.UserController{}, "post:Register")
}
