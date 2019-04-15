package controllers

import (
	"../common/dto"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UserController struct {
	beego.Controller
}

func init() {
	//todoxcs mysql 连接字符串放入配置文件中读取
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/youtubeImitation?charset=utf8")
}

func (c *UserController) Register() {
	userName := c.GetString("userName")
	email := c.GetString("email")
	phoneNumber := c.GetString("phoneNumber")
	password := c.GetString("password")
	sex, _ := c.GetInt("sex")
	o := orm.NewOrm()
	//todoxcs 修改成 beego 中 module 的方式
	if userName != "" && email != "" && phoneNumber != "" && password != "" {
		res, err := o.Raw("INSERT INTO `User`(`uid`, `userName`, `email`, `phoneNumber`, `password`, `sex`) VALUES (?, ?, ?, ?, ?, ?)", nil, userName, email, phoneNumber, password, sex).Exec()
		if err == nil {

			num, _ := res.RowsAffected()
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
