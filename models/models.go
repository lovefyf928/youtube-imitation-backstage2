package models

import (
	"../common/dto"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
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
	res, err := o.Raw(sql, args).Exec()
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


func SqlI(sql string, msg string, data interface{}, args ...interface{}) *dto.ResponseDto {
	o := orm.NewOrm()
	res, err := o.Raw(sql, nil,args).Exec()
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


func SqlS(sql string, args ...interface{})  ([]orm.Params, bool) {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw(sql, args).Values(&maps)
	if err == nil && num > 0 {
		return maps, true
	} else {
		return nil, false
	}
}
