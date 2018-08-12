package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 初始化数据库连接
func init() {
	dbHost := beego.AppConfig.String("dbHost")
	dbPort := beego.AppConfig.String("dbPort")
	dbUser := beego.AppConfig.String("dbUser")
	dbPassword := beego.AppConfig.String("dbPassword")
	dbDatabase := beego.AppConfig.String("dbDatabase")
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		dbUser, dbPassword, dbHost, dbPort, dbDatabase)
	orm.RegisterDataBase("default", "mysql", datasource, 30, 30)
	orm.RegisterModel(new(Url))
	orm.RunSyncdb("default", false, false)

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
}
