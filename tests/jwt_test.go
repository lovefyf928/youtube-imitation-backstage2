package test

import (
	"../common/authorization"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

func Test_BuildUserToken(t *testing.T) {

	secretKey := "peter"
	issuer := "peter"

	claims, err := authorization.BuildUserToken([]byte(secretKey), issuer, 1)

	if err != nil {
		t.Error(err)
	}

	t.Log(claims)
}

func Test_ParseUserToken(t *testing.T) {
	secretKey := "youtube_imitation"
	issuer := "peter"

	claims, err := authorization.BuildUserToken([]byte(secretKey), issuer, 1)

	if err != nil {
		t.Error(err)
	}

	t.Log(claims)

	orignalClaims, err := authorization.ParseUserToken(claims, []byte(secretKey))

	if err != nil {
		t.Error(err)
	}

	t.Log(orignalClaims.(jwt.MapClaims)["uid"])

}
