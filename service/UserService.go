package service

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"../common/authorization"
	"../repository"
)


//todo uid ,newPassword 参数换成 Models中的 User
func ChangeI(uid interface{}, newPassword interface{}, sex interface{}, year string, month string, day string, userName string) bool {
	date := year + "-" + month + "-" + day
	return repository.ChangeInformationRepository(uid, newPassword, sex, userName, date)
}

//todo 将Sql操作封装到 Repository 中去		(done)
func SelectUserInformation(uid interface{}) ([]orm.Params, bool) {
	return repository.SelectUserInformationRepository(uid)
}

func RegisterService(userName string, email string, phoneNumber string, password string) bool {
	return repository.RegisterRepository(userName, email, phoneNumber, password)
}

func LoginService(userName string, password string) (string, bool) {
	res, ok := repository.LoginRepository(userName, password)
	if ok {
		key := []byte(beego.AppConfig.String(authorization.TOKEN_CONFIG_NAME))
		uid, _ := strconv.Atoi(res.(string))
		token, err := authorization.BuildUserToken(key, userName, uint(uid))
		if err == nil {
			return token, true
		} else {
			return "Generating token failed", false
		}
	} else {
		return "your userName or password error", false
	}
}

func SelectUserNameService (phoneNumber string, email string) ([]string, bool) {
	return repository.SelectUserNameRepository(phoneNumber, email)
}
