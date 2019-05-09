package routers

import (
	"../controllers"
	_ "../filters"
	"github.com/astaxie/beego"
)



func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.UserController{}, "post:Register")
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/changeinformation", &controllers.UserController{}, "post:ChangeInformation")
	beego.Router("/logout", &controllers.UserController{}, "post:Logout")
	beego.Router("/selectusername", &controllers.UserController{}, "post:SelectUserName")
	beego.Router("/testToken", &controllers.UserController{}, "post:TokenSelectUsernameAndEmail")
	beego.Router("/getinformation", &controllers.UserController{}, "post:GetInformation")
}
