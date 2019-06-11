package repository

import (
	"github.com/astaxie/beego/orm"
	"time"
	"youtube-imitation-backstage2/models"
)

func UpLoadVideoRepository(uid interface{}, title string, explain string, label string, videoPath string, videoImgPath string) bool {
	maps, ok := models.SqlS("select id from channel where uid=?", uid)
	if ok {
		cid := maps[0]["id"]
		Nowtime := time.Now().UnixNano()
		arr := []interface{}{nil ,cid, nil, label, explain, title, 0, 0, 0, Nowtime, videoPath, videoImgPath}
		ok := models.SqlI1("INSERT INTO `video`(`id`, `channelId`, `channelClassification`, `category`, `videoIntroduction`, `name`, `viewCount`, `good`, `bad`, `releaseTime`, `videoPath`, `videoImg`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)", arr)
		if ok {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func IndexRenderRepository() ([]orm.Params, bool) {
	maps, ok := models.SqlS("select id, name, subscriber from channel order by subscriber DESC")
	if ok {
		for key, _ := range maps{
			maps2, ok := models.SqlS("select * from video where channelId=?", maps[key]["id"])
			if ok {
				maps[key]["video"] = maps2
			}
		}
		return maps, ok
	} else {
		return nil, false
	}
}

func SearchVideoRepository(keyWord string) ([]orm.Params, bool) {
	maps, ok := models.SqlS("select * from video where name like '%"+ keyWord + "%' order by viewCount desc")
	if ok {
		return maps, ok
	} else {
		return nil, ok
	}
}

func SubscribeRepository(userId interface{}, creatorId interface{}) []orm.Params{
	_, ok := models.SqlS("select * from subscribe where userId=? and Subscribed=?", userId, creatorId)
	if !ok {
		ok := models.SqlI1("INSERT INTO `subscribe`(`userId`, `Subscribed`) VALUES (?,?)", []interface{}{userId, creatorId})
		if ok {
			maps, ok1 := models.SqlS("select subscriber from channel where id=?", creatorId)
			if ok1 {
				v, _ := maps[0]["subscriber"].(int)
				ok2 := models.SqlI1("UPDATE `channel` SET `subscriber`=? WHERE id=?", []interface{}{v + 1, creatorId})
				if ok2 {
					maps1, ok3 := models.SqlS("select subscriber from channel where id=?", creatorId)
					if ok3 {
						return maps1
					}
				}
			}
		}
	} else {
		ok := models.SqlIDU("DELETE FROM `subscribe` WHERE userId=? and Subscribed=?","", nil, userId, creatorId)
		if ok {
			maps, ok1 := models.SqlS("select subscriber from channel where id=?", creatorId)
			if ok1 {
				v, _ := maps[0]["subscriber"].(int)
				ok2 := models.SqlI1("UPDATE `channel` SET `subscriber`=? WHERE id=?", []interface{}{v - 1, creatorId})
				if ok2 {
					maps1, ok3 := models.SqlS("select subscriber from channel where id=?", creatorId)
					if ok3 {
						return maps1
					}
				}
			}
		}
	}
	return nil
}

func SubscribeListRepository(uid interface{}) []orm.Params {
	maps, ok := models.SqlS("select name FROM channel WHERE id=(SELECT subscribed from subscribe WHERE userId=?)", uid)
	if ok {
		return maps
	} else {
		return nil
	}
}

func ViewCountRepository(videoId interface{}) bool {
	maps, ok := models.SqlS("select viewCount from video where id=?", videoId)
	if ok {
		v, _ := maps[0]["viewCount"].(int)
		return models.SqlI1("UPDATE `video` SET `viewCount`=? WHERE id=?", []interface{}{v + 1, videoId})
	} else {
		return false
	}
}
