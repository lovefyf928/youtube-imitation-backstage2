package repository

import (
	"github.com/astaxie/beego/orm"
	"../models"
)

func RegisterRepository(userName string, email string, phoneNumber string, password string) bool {
	if models.SqlI("INSERT INTO `User`(`uid`, `userName`, `email`, `phoneNumber`, `password`) VALUES (?, ?, ?, ?, md5(?))", nil, userName, email, phoneNumber, password) {
		return true
	} else {
		return false
	}
}

func LoginRepository(userName string, password string) (interface{}, bool) {
	maps, ok := models.SqlS("select uid from user where userName = ? and password = md5(?)", userName, password)
	if ok {
		return maps[0]["uid"], true
	} else {
		return nil, false
	}
}

func ChangeInformationRepository(uid interface{}, newPassword interface{}, sex interface{}, userName string, date string) bool {
	if newPassword == "" {
		maps, ok := models.SqlS("select password from user where uid=?", uid)
		if ok {
			newPassword = maps[0]["password"]
			return models.SqlIDU("UPDATE `user` SET `userName`=?,`password`=?,`sex`=?,`birthday`=? WHERE uid=?", "update successful", nil, userName, newPassword, sex, date, uid)
		}
	}
	return models.SqlIDU("UPDATE `user` SET `userName`=?,`password`=md5(?),`sex`=?,`birthday`=? WHERE uid=?", "update successful", nil, userName, newPassword, sex, date, uid)
}

func SelectUserNameRepository(phoneNumber string, email string) ([]string, bool) {
	maps, ok := models.SqlS("select userName,email from user where phoneNumber=? or email=?", phoneNumber, email)
	if ok {
		arr := []string{maps[0]["userName"].(string), maps[0]["email"].(string)}
		return arr, true
	} else {
		return []string{}, false
	}
}

func SelectUserInformationRepository(uid interface{}) ([]orm.Params, bool) {
	maps, ok := models.SqlS("select * from user where uid=?", uid)
	if ok {
		return maps, true
	} else {
		return nil, false
	}
}
