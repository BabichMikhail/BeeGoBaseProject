package main

import (
	"./console"
	"./middlewares"
	"./models"
	_ "./routers"
	"./utils"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/astaxie/beego"
	"github.com/beego/compress"
	"github.com/beego/wetalk/setting"
)

func SettingCompress() {
	configuration, err := compress.LoadJsonConf("conf/compress.json", true, "")
	utils.CheckError(err)

	configuration.RunCommand()
	configuration.RunCompress(true, false, true)
	utils.CheckError(beego.AddFuncMap("compress_js", configuration.Js.CompressJs))
	utils.CheckError(beego.AddFuncMap("compress_css", configuration.Css.CompressCss))
}

func init() {
	beego.SetStaticPath("/images", "static/images")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")

	middlewares.InitMiddleware()

	Cfg, _ := goconfig.LoadConfigFile("conf/app.conf")
	Cfg.BlockMode = false

	beego.BConfig.WebConfig.EnableXSRF = false

	driverName := Cfg.MustValue("orm", "driver_name", "")
	dataSource := Cfg.MustValue("orm", "data_source", "")
	maxIdle := Cfg.MustInt("orm", "max_idle_conn", 30)
	maxOpen := Cfg.MustInt("orm", "max_open_conn", 50)

	models.RegisterDatabase(driverName, dataSource, maxIdle, maxOpen)

	setting.SecretKey = Cfg.MustValue("app", "secret_key")
	if len(setting.SecretKey) == 0 {
		fmt.Println("Please set your secret_key in app.conf file")
	}

	setting.LoginRememberDays = Cfg.MustInt("app", "login_remember_days", 7)
	setting.LoginMaxRetries = Cfg.MustInt("app", "login_max_retries", 5)
	setting.LoginFailedBlocks = Cfg.MustInt("app", "login_failed_blocks", 10)

	setting.CookieRememberName = Cfg.MustValue("app", "cookie_remember_name", "")
	setting.CookieUserName = Cfg.MustValue("app", "cookie_user_name", "")
}

func main() {
	// TODO web module
	// TODO console module
	if console.Console {
		fmt.Println("Running as console application")
		console.Run()
	} else {
		fmt.Println("Running as web application")
		SettingCompress()
		beego.Run()
	}
}
