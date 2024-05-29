package main

import (
	"net/url"

	_ "github.com/udistrital/inscripcion_crud/routers"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/customerror"

	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
	"github.com/udistrital/auditoria"
	"github.com/udistrital/utils_oas/xray"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://"+
		beego.AppConfig.String("PGuser")+":"+
		url.QueryEscape(beego.AppConfig.String("PGpass"))+"@"+
		beego.AppConfig.String("PGurls")+":"+
		beego.AppConfig.String("PGport")+"/"+
		beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+
		beego.AppConfig.String("PGschemas")+"")

	if beego.BConfig.RunMode == "dev" {
		/*
			// Database alias.
			name := "default"

			// Drop table and re-create.
			force := false

			// Print log.
			verbose := true

			// Error.
			err := orm.RunSyncdb(name, force, verbose)
			if err != nil {
				fmt.Println(err)
			}
		*/
	}
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		orm.Debug = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// logPath := "{\"filename\":\""
	// logPath += beego.AppConfig.String("logPath")
	// logPath += "\"}"
	// logs.SetLogger(logs.AdapterFile, logPath)

	beego.ErrorController(&customerror.CustomErrorController{})
	xray.InitXRay()
	apistatus.Init()
	auditoria.InitMiddleware()
	beego.Run()
}
