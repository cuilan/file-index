package main

import (
	"fmt"
	"github.com/cuilan/file-index/configer"
	"github.com/cuilan/file-index/service"
)

func init() {
	if err := configer.InitLogger(); err != nil {
		fmt.Println("init logger err", err)
	}
	configer.InitConfig()
}

func main() {
	service.Index()
}
