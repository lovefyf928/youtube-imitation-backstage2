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
	beego.Router("/changeinformation", &controllers.UserController{}, "post:ChangeInformation")
	beego.Router("/logout", &controllers.UserController{}, "post:Logout")
	beego.Router("/selectusername", &controllers.UserController{}, "post:SelectUserName")
	beego.Router("/testToken", &controllers.UserController{}, "post:TokenSelectUsernameAndEmail")
	beego.Router("/getinformation", &controllers.UserController{}, "post:GetInformation")
	beego.Router("/uploadvideo", &controllers.VideoController{}, "post:UpLoadVideo")
	beego.Router("/indexrender", &controllers.VideoController{}, "get:IndexRender")
	beego.Router("/searchvideo", &controllers.VideoController{}, "get:SearchVideo")
	beego.Router("/subscribe", &controllers.VideoController{}, "post:Subscribe")
	beego.Router("/subscribelist", &controllers.VideoController{}, "post:SubscribeList")
	beego.Router("/viewcount", &controllers.VideoController{}, "post:ViewCount")
}
