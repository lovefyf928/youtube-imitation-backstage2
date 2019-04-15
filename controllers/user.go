package controllers

import (
	"youtube-imitation-backstage2/common/dto"
	"youtube-imitation-backstage2/models"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}


func (c *UserController) Register() {
	userName := c.GetString("userName")
	email := c.GetString("email")
	phoneNumber := c.GetString("phoneNumber")
	password := c.GetString("password")
	sex, _ := c.GetInt("sex")
	if userName != "" && email != "" && phoneNumber != "" && password != "" {
		c.Data["json"] = models.SqlIDU("INSERT INTO `User`(`uid`, `userName`, `email`, `phoneNumber`, `password`, `sex`) VALUES (?, ?, ?, ?, ?, ?)", "Register user success", nil, nil, userName, email, phoneNumber, password, sex)
		c.ServeJSON()
		return
	}
	c.Data["json"] = dto.NewResponseDto(false, dto.FORBBDIEN, "plz check your parameter", nil)
	c.ServeJSON()
}

func (c *UserController) login() {

}
