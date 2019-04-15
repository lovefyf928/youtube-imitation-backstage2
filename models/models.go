package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"youtube-imitation-backstage2/common/dto"
)

func init()  {

	dbname := beego.AppConfig.String("dbName")
	dbdname := beego.AppConfig.String("dbdName")
	ds := beego.AppConfig.String("dataSource")

	orm.RegisterDriver(dbdname, orm.DRMySQL)

	orm.RegisterDataBase(dbname, dbdname, ds)
}

func SqlIDU(sql string, msg string, data interface{}, args ...interface{}) *dto.ResponseDto {
	o := orm.NewOrm()
	res, err := o.Raw(sql, nil, args).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num > 0 {
			return dto.NewResponseDto(true, dto.SUCCESS, msg, data)
		} else {
			return dto.NewResponseDto(false, dto.FORBBDIEN, "plz check your parameter", nil)
		}
	}
	return dto.NewResponseDto(false, dto.FORBBDIEN, "plz check your parameter", nil)
}
