package controllers

import (
	"../common/authorization"
	"../common/dto"
	"../models"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

type UserController struct {
	beego.Controller
}



func (c *UserController) Register() {
	userName := c.GetString("userName")
	email := c.GetString("email")
	phoneNumber := c.GetString("phoneNumber")
	password := c.GetString("password")
	beego.Info(password)
	if userName != "" && email != "" && phoneNumber != "" && password != "" {
		c.Data["json"] = models.SqlI("INSERT INTO `User`(`uid`, `userName`, `email`, `phoneNumber`, `password`) VALUES (?, ?, ?, ?, ?)", "Register user success", nil, nil, userName, email, phoneNumber, password)
		c.ServeJSON()
		return
	}
	c.Data["json"] = dto.NewResponseDto(false, dto.FORBBDIEN, "plz check your parameter", nil)
	c.ServeJSON()
}

func (c *UserController) Login() {
	userName := c.GetString("userName")
	password := c.GetString("password")
	if c.GetSession(userName) == "login" {
		c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("you are already login")
		c.ServeJSON()
		return
	}
	if userName != "" && password != "" {
		maps, ok := models.SqlS("select uid from user where userName = ? and password = ?", userName, password)
		if ok {
			key := []byte(beego.AppConfig.String(authorization.TOKEN_CONFIG_NAME))
			uid, _ := strconv.Atoi(maps[0]["uid"].(string))
			token, err := authorization.BuildUserToken(key, userName, uint(uid))
			if err == nil {
				c.SetSession(userName, "login")
				c.Data["json"] = dto.NewSuccessResponseDto(map[string]interface{}{"msg": "login success", "token": token})
			} else {
				c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("Generating token failed")
			}
		} else {
			c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("Error in your userName or password")
		}
		c.ServeJSON()
	}
}

func (c *UserController) ChangePassword() {
	oldPassword := c.GetString("oldPassword")
	newPassword := c.GetString("newPassword")
	if oldPassword != "" && newPassword != "" {
		var token = c.Ctx.Request.Header[authorization.TOKEN_HEADER_NAME]
		userClaims, _ := authorization.ParseUserToken(token[0], []byte(beego.AppConfig.String(authorization.TOKEN_CONFIG_NAME)))
		uid := userClaims.(jwt.MapClaims)["uid"]
		maps, ok := models.SqlS("select password from user where uid=?", uid)
		if ok {
			if oldPassword == maps[0]["password"] {
				c.Data["json"] = models.SqlIDU("UPDATE `user` SET `password`=? WHERE uid=?", "update password successful", nil, newPassword, uid)
			} else {
				c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("your old password error")
			}
		}
	} else {
		c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("plz enter your old password and new password")
	}
	c.ServeJSON()
}

func (c *UserController) Logout() {
	userName := c.GetString("userName")
	str := c.GetSession(userName)
	if str == nil {
		c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("not logged on")
		c.ServeJSON()
		return
	}
	if userName != ""  {
		c.DelSession(userName)
		c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("Logout success")
	} else {
		c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("Please introduce username")
	}
	c.ServeJSON()
}

func (c *UserController) SelectUserName() {
	phoneNumber := c.GetString("phoneNumber")
	email := c.GetString("email")
	beego.Info(phoneNumber)
	maps, ok := models.SqlS("select userName,email from user where phoneNumber=? or email=?", phoneNumber, email)
	if ok {
		username := maps[0]["userName"].(string)
		email := maps[0]["email"].(string)
		beego.Info(username)
		c.Data["json"] = dto.NewSuccessResponseDto(map[string]interface{}{"userName": username, "email": email})
	} else {
		c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("your phone number or email error")
	}
	c.ServeJSON()
}
