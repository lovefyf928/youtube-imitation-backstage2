package service

import (
	"github.com/astaxie/beego/orm"
	"youtube-imitation-backstage2/repository"
)

func UpLoadVideoService(uid interface{}, title string, explain string, label string, videoPath string, videoImgPath string) bool {
	if repository.UpLoadVideoRepository(uid, title, explain, label, videoPath, videoImgPath) {
		return true
	} else {
		return false
	}
}

func IndexRenderService() ([]orm.Params, bool) {
	return repository.IndexRenderRepository()
}

func SearchVideoService(keyWord string) ([]orm.Params, bool) {
	return repository.SearchVideoRepository(keyWord)
}


func SubscribeService(userId interface{}, creatorId interface{}) []orm.Params{
	return repository.SubscribeRepository(userId, creatorId)
}

func SubscribeListService(userId interface{}) []orm.Params {
	return repository.SubscribeListRepository(userId)
}

func ViewCountService(videoId interface{}) bool {
	return repository.ViewCountRepository(videoId)
}
