package main

import (
	_ "sstest/todoapi/docs"
	_ "sstest/todoapi/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbAlias := beego.AppConfig.String("db::dbAlias")
	dbName := beego.AppConfig.String("db::dbName")
	dbDriver := beego.AppConfig.String("db::dbDriver")
	dbHost := beego.AppConfig.String("db::dbHost")
	dbPort := beego.AppConfig.String("db::dbPort")
	dbProtocol := beego.AppConfig.String("db::dbProtocol")
	dbUser := beego.AppConfig.String("db::dbUser")
	dbPwd := beego.AppConfig.String("db::dbPwd")

	var dbURL string

	if dbDriver == "postgres" {
		dbURL = dbProtocol + "://" + dbUser + ":" + dbPwd + "@" + dbHost + ":" + dbPort + "/" + dbName
	} else if dbDriver == "mysql" {
		dbURL = dbUser + ":" + dbPwd + "@" + dbProtocol + "(" + dbHost + ":" + dbPort + ")" + "/" + dbName + "?charset=utf8"
	}
	orm.RegisterDataBase(dbAlias, dbDriver, dbURL)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		beego.BConfig.WebConfig.StaticDir["/s1"] = "static1"
	}
	//orm.RunCommand()
	beego.BConfig.Listen.EnableAdmin = true
	beego.Run()
}

