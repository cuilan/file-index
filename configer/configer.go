package configer

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/cuilan/file-index/orm"
	"os"
)

var (
	AppName     string
	ScanPath    string
	StorageMode string
)

func InitConfig() {
	printBanner()
	fmt.Println("======== init system config start ========>")
	AppName = beego.AppConfig.String("app_name")
	ScanPath = beego.AppConfig.String("scan_path")
	if ScanPath == "" {
		fmt.Println("scan_path is not configured")
		os.Exit(1)
	}
	StorageMode = beego.AppConfig.String("storage_mode")
	if StorageMode == "" {
		fmt.Println("storage_mode is not configured")
		os.Exit(1)
	}
	fmt.Println("<======== init system config end ========")
	initDb()
}

func initDb() {
	logs.Info("load database configure and start orm begin")
	var err error
	if StorageMode == "mysql" {
		_, err = beego.AppConfig.GetSection("mysql")
		if err != nil {
			logs.Error("init mysql connect error: %s", err.Error())
			os.Exit(1)
		}
		host := beego.AppConfig.String("mysql::host")
		user := beego.AppConfig.String("mysql::user")
		password := beego.AppConfig.String("mysql::password")
		port, _ := beego.AppConfig.Int("mysql::port")
		db := beego.AppConfig.String("mysql::db")
		charset := beego.AppConfig.String("mysql::charset")
		maxActive, _ := beego.AppConfig.Int("mysql::maxactive")
		maxIdle, _ := beego.AppConfig.Int("mysql::maxidle")
		orm.StartMysql(host, user, password, db, charset, maxActive, maxIdle, port)
	} else if StorageMode == "sqlite" {
		_, err = beego.AppConfig.GetSection("sqlite")
		sqlitePath := beego.AppConfig.String("sqlite::sqlite_path")
		orm.StartSqlite(sqlitePath)
	}
	logs.Info("load database configure and start orm end")
}

func printBanner() {
	fmt.Println("  __ _ _         _           _           ")
	fmt.Println(" / _(_) |       (_)         | |          ")
	fmt.Println("| |_ _| | ___    _ _ __   __| | _____  __")
	fmt.Println("|  _| | |/ _ \\  | | '_ \\ / _` |/ _ \\ \\/ /")
	fmt.Println("| | | | |  __/  | | | | | (_| |  __/>  < ")
	fmt.Println("|_| |_|_|\\___|  |_|_| |_|\\__,_|\\___/_/\\_\\")
	fmt.Println("")
}
