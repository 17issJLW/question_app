package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
	_ "test/routers"

	"github.com/astaxie/beego"
)

func init()  {
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlurls := beego.AppConfig.String("mysqlurls")
	mysqldb := beego.AppConfig.String("mysqldb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:jlw123456@tcp(211.159.186.47:3306)/app?charset=utf8")
	orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+")/"+mysqldb+"?charset=utf8")
	orm.RunSyncdb("default",false,true)
}

func main() {
	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//beego.InsertFilter("*", beego.BeforeRouter, auth.Basic("appdoc","appdoc123456"))
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.SetStaticPath("/down", "./static")
	beego.BConfig.Log.AccessLogs = true
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.Run()
}
