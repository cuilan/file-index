package configer

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

func InitLogger() (err error) {
	BConfig, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		fmt.Println("config init error:", err)
		return
	}
	logConf := make(map[string]interface{})
	logConf["filename"] = BConfig.String("log::filename")

	maxlines, lerr := BConfig.Int64("log::maxlines")
	if lerr != nil {
		maxlines = 1000
	}
	logConf["maxlines"] = maxlines

	maxsize, serr := BConfig.Int64("log::maxsize")
	if serr != nil {
		maxsize = 1024
	}
	logConf["maxsize"] = maxsize

	daily, _ := BConfig.Bool("log::daily")
	logConf["daily"] = daily

	maxdays, _ := BConfig.Int("log::maxdays")
	logConf["maxdays"] = maxdays

	rotate, _ := BConfig.Bool("log::rotate")
	logConf["rotate"] = rotate

	level, _ := BConfig.Int("log::level")
	logConf["level"] = level

	confStr, err := json.Marshal(logConf)
	if err != nil {
		fmt.Println("marshal failed,err:", err)
		return
	}
	_ = logs.SetLogger(logs.AdapterFile, string(confStr))
	logs.SetLogFuncCall(true)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	return
}
