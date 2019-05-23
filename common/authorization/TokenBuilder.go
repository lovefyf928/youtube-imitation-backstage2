package authorization

import (
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TOKEN_HEADER_NAME = "Authorization"
const TOKEN_CONFIG_NAME = "secretkey"

type UserClaims struct {
	jwt.StandardClaims

	UserId uint `json:"uid"`
}

func BuildUserToken(SecretKey []byte, issuer string, Uid uint) (tokenString string, err error) {
	claims := &UserClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
			Issuer:    issuer,
		},
		Uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(SecretKey)
	return
}

func ParseUserToken(tokenSrt string, SecretKey []byte) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims

	return
}

func GetTokenValueByKey(token string,key string) interface{} {
	userClaims, _ := ParseUserToken(token, []byte(beego.AppConfig.String(TOKEN_CONFIG_NAME)))
	res := userClaims.(jwt.MapClaims)[key]
	return res
}
