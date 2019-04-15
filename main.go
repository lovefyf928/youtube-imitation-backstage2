package main

import (
	"./common/authorization"
	"./controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func main() {
	//todo router 放入单独文件 使用 import
	beego.Router("/", &controllers.MainController{})

	beego.Router("/register", &controllers.UserController{}, "post:Register")

	//todo filter 放入单独文件 使用 import
	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {

		var token = ctx.Request.Header[authorization.TOKEN_HEADER_NAME]

		if token == nil || len(token) == 0 {
			ctx.Redirect(http.StatusUnauthorized, "/")
			return
		}

		userClaims, err := authorization.ParseUserToken(token[0], []byte(beego.AppConfig.String(authorization.TOKEN_CONFIG_NAME)))

		if err != nil {
			ctx.Redirect(http.StatusUnauthorized, "/")
			return
		}

		userId := userClaims.(jwt.MapClaims)["uid"]

		if userId == nil {
			ctx.Redirect(http.StatusUnauthorized, "/")
			return
		}
	})

	beego.Run()

}
