package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init()  {
	//beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
	//
	//	var token = ctx.Request.Header[authorization.TOKEN_HEADER_NAME]
	//
	//	if token == nil || len(token) == 0 {
	//		ctx.Redirect(http.StatusUnauthorized, "/")
	//		return
	//	}
	//
	//	userClaims, err := authorization.ParseUserToken(token[0], []byte(beego.AppConfig.String(authorization.TOKEN_CONFIG_NAME)))
	//
	//	if err != nil {
	//		ctx.Redirect(http.StatusUnauthorized, "/")
	//		return
	//	}
	//
	//	userId := userClaims.(jwt.MapClaims)["uid"]
	//
	//	if userId == nil {
	//		ctx.Redirect(http.StatusUnauthorized, "/")
	//		return
	//	}
	//})
	//
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

}
