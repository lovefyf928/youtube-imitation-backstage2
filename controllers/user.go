package controllers

import (
	"regexp"
	"../common/authorization"
	"../common/dto"
	"../service"
	"github.com/astaxie/beego"
	DtoUser "../controllers/dto"
)

type GetUser struct {
	 UserName string `form:"userName"`
	 Email string `form:"email"`
	 PhoneNumber string `form:"phoneNumber"`
	 PassWord string `form:"password"`
	 Sex interface{} `form:"sex"`
	 Year string `form:"year"`
	 Month string `form:"month"`
	 Day string `form:"day"`
	 NewPassword string `form:"newPassword"`
}


type UserController struct {
	beego.Controller
}

//todo 查询beego的文档，是否是这样获取参数，能否通过 自动组装 dto的方式，参看 controllers 文件夹下 description.md 中的伪代码	(done)

func (c *UserController) Register() {
	u := GetUser{}
	c.ParseForm(&u)
	userName := u.UserName
	email := u.Email
	phoneNumber := u.PhoneNumber
	password := u.PassWord
	EmailMatched, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, email)
	phoneNumberMatched, _ := regexp.MatchString(`^1[34578]\d{9}$`, phoneNumber)
	if userName != "" && EmailMatched && phoneNumberMatched && len(password) >= 8 {
		if service.RegisterService(userName, email, phoneNumber, password) {
			c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("register user success")
		} else {
			c.Data["json"] = dto.NewResponseDto(false, dto.FORBBDIEN, "register user error", nil)
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = dto.NewResponseDto(false, dto.FORBBDIEN, "your params format error", nil)
	c.ServeJSON()
}

func (c *UserController) Login() {
	u := GetUser{}
	c.ParseForm(&u)
	userName := u.UserName
	password := u.PassWord
	if c.GetSession(userName) == "login" {
		c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("you are already login")
		c.ServeJSON()
		return
	}
	if userName != "" && password != "" {
		res, ok := service.LoginService(userName, password)
		if ok {
			c.Data["json"] = dto.NewResponseDto(true, dto.SUCCESS, "login success", DtoUser.LoginDto(res))
		} else {
			c.Data["json"] = dto.NewResponseDto(false, dto.FORBBDIEN, res, nil)
		}
		c.ServeJSON()
	}
}

func (c *UserController) ChangeInformation() {
	u := GetUser{}
	c.ParseForm(&u)
	sex := u.Sex
	year := u.Year
	month := u.Month
	day := u.Day
	userName := u.UserName
	newPassword := u.NewPassword
	if newPassword != "" && len(newPassword) < 8 {
		c.Data["json"] = dto.NewResponseDto(false, dto.FORBBDIEN, "your new password format error", nil)
		c.ServeJSON()
		return
	}
	if sex != nil && year != "" && month != "" && day != "" && userName != "" {
		var token = c.Ctx.Request.Header[authorization.TOKEN_HEADER_NAME]
		//todo authorization中封装一个方法，提供key 获取value，例如  uid:=authorization.getTokenValueByKey("uid");	(done)
		uid := authorization.GetTokenValueByKey(token[0], "uid")
		res := service.ChangeI(uid, newPassword, sex, year, month, day, userName)
		if res {
			c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("update successful")
		} else {
			c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("error")
		}
		c.ServeJSON()
	}
}

func (c *UserController) Logout() {
	u := GetUser{}
	c.ParseForm(&u)
	userName := u.UserName
	str := c.GetSession(userName)
	if str == nil {
		c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("not logged on")
		c.ServeJSON()
		return
	}
	if userName != "" {
		c.DelSession(userName)
		c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("Logout success")
	} else {
		c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("Please introduce username")
	}
	c.ServeJSON()
}

func (c *UserController) SelectUserName() {
	u := GetUser{}
	c.ParseForm(&u)
	phoneNumber := u.PhoneNumber
	email := u.Email
	arr, ok := service.SelectUserNameService(phoneNumber, email)
	if ok {
		username := arr[0]
		email := arr[1]
		c.Data["json"] = dto.NewSuccessResponseDto(DtoUser.TokenSelectUsernameAndEmailDto(username, email))
	} else {
		c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("your phone number or email error")
	}
	c.ServeJSON()
}

func (c *UserController) TokenSelectUsernameAndEmail() {
	var token = c.Ctx.Request.Header[authorization.TOKEN_HEADER_NAME]
	uid := authorization.GetTokenValueByKey(token[0], "uid")
	arr, ok := service.SelectUserInformation(uid)
	if ok {
		c.Data["json"] = dto.NewSuccessResponseDto(DtoUser.TokenSelectUsernameAndEmailDto(arr[0]["userName"], arr[0]["email"]))
	} else {
		c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("your token error")
	}
	c.ServeJSON()
}

func (c *UserController) GetInformation() {
	var token = c.Ctx.Request.Header[authorization.TOKEN_HEADER_NAME]
	uid := authorization.GetTokenValueByKey(token[0], "uid")
	arr, ok := service.SelectUserInformation(uid)
	if ok {
		c.Data["json"] = dto.NewSuccessResponseDto(DtoUser.Ud(arr[0]["uid"], arr[0]["userName"], arr[0]["email"], arr[0]["phoneNumber"], nil, arr[0]["sex"], arr[0]["birthday"], arr[0]["code"]))

	} else {
		c.Data["json"] = dto.NewSuccessResponseDtoNilMsg("error")
	}
	c.ServeJSON()
}

