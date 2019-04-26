package filters

import (
	"../common/authorization"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func init()  {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))


	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {

		//if ctx.Request.RequestURI != "/register" && ctx.Request.RequestURI != "/login" && ctx.Request.RequestURI != "/selectusername" {

			var token= ctx.Request.Header[authorization.TOKEN_HEADER_NAME]

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
		//}
	})

}
