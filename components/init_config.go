package components

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	AppName  string
	ScanPath string
)

func InitConfiger() {
	printBanner()
	logs.Info("init app config start...")
	AppName = beego.AppConfig.String("app::app_name")
	ScanPath = beego.AppConfig.String("app::scan_path")
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
