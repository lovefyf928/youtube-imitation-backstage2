package controllers

import (
	"github.com/astaxie/beego"
	"youtube-imitation-backstage2/common/authorization"
	"youtube-imitation-backstage2/common/dto"
	"youtube-imitation-backstage2/service"
)

type VideoController struct {
	beego.Controller
}

type video struct {
	Title string `form:"title"`
	Explain string `form:"explain"`
	Label string `form:"label"`
	Token string `form:"token"`
	KeyWold string `form:"keywold"`
	CreatorId interface{} `form:"creatorid"` //传channel id
	VideoId interface{} `form:"videoid"`
}

func (c *VideoController) Save(fileKey string, filePath string) (string, bool) {
	v, h, err := c.GetFile(fileKey)
	beego.Info(err)
	if err != nil {
		return  "", false
	}
	defer v.Close()
	e := c.SaveToFile(fileKey, filePath + h.Filename)
	if e != nil {
		return  "", false
	} else {
		return  filePath + h.Filename, true
	}
}


func (c *VideoController) UpLoadVideo(){
	RequestData := video{}
	c.ParseForm(&RequestData)
	beego.Info(RequestData.Label)
	if RequestData.Label != "" && RequestData.Explain != "" && RequestData.Title != "" && RequestData.Token != "" {
		uid := authorization.GetTokenValueByKey(RequestData.Token, "uid")
		videoPath := "static/video/"
		videoImgPath := "static/img/videoImg/"
		FileName, e := c.Save("video", videoPath)
		FileName2, err := c.Save("videoImg", videoImgPath)
		beego.Info(err)
		ok := service.UpLoadVideoService(uid, RequestData.Title, RequestData.Explain, RequestData.Label, FileName, FileName2)
		if e && err && ok {
			c.Data["json"] = dto.NewResponseDto(true, dto.SUCCESS, "save file success", nil)
		} else {
			c.Data["json"] = dto.NewResponseDto(false, dto.INTERNATL_ERROR, "save file error", nil)
		}
		c.ServeJSON()
	}
}

func (c *VideoController) IndexRender() {
	token := c.Ctx.Request.Header[authorization.TOKEN_HEADER_NAME]
	if token == nil || len(token) == 0 {
		maps, ok := service.IndexRenderService()
		if ok {
			c.Data["json"] = maps
			c.ServeJSON()
		}
	}
}

func (c *VideoController) SearchVideo() {
	data := video{}
	c.ParseForm(&data)
	if data.KeyWold != "" {
		maps, _ := service.SearchVideoService(data.KeyWold)
		c.Data["json"] = maps
		c.ServeJSON()
	}
}

func (c *VideoController) Subscribe() {
	token := c.Ctx.Request.Header[authorization.TOKEN_HEADER_NAME]
	uid := authorization.GetTokenValueByKey(token[0], "uid")
	data := video{}
	c.ParseForm(&data)
	if data.CreatorId != nil {
		c.Data["json"] = service.SubscribeService(uid, data.CreatorId)
	} else {
		c.Data["json"] = dto.NewResponseDto(false, dto.FORBBDIEN, "not get channel id", nil)
	}
	c.ServeJSON()
}

func (c *VideoController) SubscribeList() {
	token := c.Ctx.Request.Header[authorization.TOKEN_HEADER_NAME]
	uid := authorization.GetTokenValueByKey(token[0], "uid")
	c.Data["json"] = service.SubscribeListService(uid)
	c.ServeJSON()
}

func (c *VideoController) ViewCount() {
	data := video{}
	c.ParseForm(&data)
	if data.VideoId != nil{
		if service.ViewCountService(data.VideoId) {
			c.Data["json"] = dto.NewResponseDto(true, dto.SUCCESS, "update view count", nil)
		} else {
			c.Data["json"] = dto.NewResponseDto(false, dto.FORBBDIEN, "update view count error", nil)
		}
	} else {
		c.Data["json"] = dto.NewResponseDto(false, dto.FORBBDIEN, "update view count error", nil)
	}
	c.ServeJSON()
}
