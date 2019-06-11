package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"youtube-imitation-backstage2/common/authorization"
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


		index := strings.Index(ctx.Request.RequestURI, "?")
		url := ""
		if index != -1 {
			url = ctx.Request.RequestURI[:index]
		} else {
			url = ctx.Request.RequestURI
		}

		if url != "/register" && url != "/login" && url != "/selectusername" && url != "/uploadvideo" && url != "/indexrender" && url != "/searchvideo" {

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
		}
	})

}
