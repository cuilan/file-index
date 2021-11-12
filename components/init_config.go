package components

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/cuilan/go-common/load"
	"gorm.io/gorm"
)

var (
	AppName  string
	ScanPath string
	Db       *gorm.DB
)

func InitConfiger() {
	printBanner()
	logs.Info("init app config start...")
	AppName = beego.AppConfig.String("app::app_name")
	ScanPath = beego.AppConfig.String("app::scan_path")
	// 加载日志配置
	load.LoadLoggerConfig("./conf/log.conf")
	// 读取配置，并初始化数据源
	Db = load.LoadMysqlConfig("./conf/mysql.conf")
	logs.Info("init app config end...")
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
