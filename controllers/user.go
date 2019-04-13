package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"youtube-imitation-backstage/common/dto"
	_ "github.com/go-sql-driver/mysql"
)

type UserController struct {
	beego.Controller
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/youtubeImitation?charset=utf8")
}

func (c *UserController) Register(){
	userName := c.GetString("userName")
	email := c.GetString("email")
	phoneNumber := c.GetString("phoneNumber")
	password := c.GetString("password")
	sex, _ := c.GetInt("sex")
	o := orm.NewOrm()
	beego.Info(email)
	if userName != "" && email != "" && phoneNumber != "" && password != "" {
		res, err := o.Raw("INSERT INTO `User`(`uid`, `userName`, `email`, `phoneNumber`, `password`, `sex`) VALUES (?, ?, ?, ?, ?, ?)", nil, userName, email, phoneNumber, password, sex).Exec()
		if err == nil {
			beego.Info(111111)
			num, _:= res.RowsAffected()
			if num > 0 {
				c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("Register user success!!!")
			} else {
				c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("Register user failed!!!")
			}
			c.ServeJSON()
		}
	}
}


func (c *UserController) login() {

}
