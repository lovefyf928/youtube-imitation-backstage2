package service

import (
	"../../common/dto"
	"../../models"
)

func ChangePWD (uid interface{}, oldPassword string, newPassword string) *dto.ResponseDto {
	maps, ok := models.SqlS("select password from user where uid=?", uid)
	if ok {
		if oldPassword == maps[0]["password"] {
			return models.SqlIDU("UPDATE `user` SET `password`=? WHERE uid=?", "update password successful", nil, newPassword, uid)
		} else {
			return dto.NewSuccessResponseDtoNilMsg("your old password error")
		}
	} else {
		return dto.NewSuccessResponseDtoNilMsg("plz enter your old password and new password")
	}
}
