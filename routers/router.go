package routers

import (
	"youtube-imitation-backstage2/controllers"
	_ "youtube-imitation-backstage2/filters"
	"github.com/astaxie/beego"
)



func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.UserController{}, "post:Register")
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/changepassword", &controllers.UserController{}, "post:ChangePassword")
	beego.Router("/logout", &controllers.UserController{}, "post:Logout")
	beego.Router("/selectusername", &controllers.UserController{}, "post:SelectUserName")
}
