package routers

import (
	"github.com/astaxie/beego"
	"youtube-imitation-backstage2/controllers"
)



func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.UserController{}, "post:Register")
}
