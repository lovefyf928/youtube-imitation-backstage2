package service

import (
	"github.com/astaxie/beego/orm"
	"../../models"
)


func ChangeI (uid interface{}, newPassword interface{}, sex interface{}, year string, month string, day string, userName string) bool {
	if newPassword == "" {
		maps, ok := models.SqlS("select password from user where uid=?", uid)
		if ok {
			newPassword = maps[0]["password"]
		}
	}
	date := year + "-" + month + "-" + day
	res := models.SqlIDU("UPDATE `user` SET `userName`=?,`password`=?,`sex`=?,`birthday`=? WHERE uid=?", "update successful", nil, userName, newPassword, sex, date, uid)
	if res.Success {
		return true
	}
	return false
}

func SelectUserInformation(uid interface{}) ([]orm.Params, bool)  {
	maps, ok := models.SqlS("select * from user where uid=?", uid)
	if ok {
		return maps, true
	} else {
		return nil, false
	}
}
