package service

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/cuilan/file-index/components"
	"os"
	"path/filepath"
	"strings"
)

func Index() {
	logs.Info("Application [%s] started...", components.AppName)
	logs.Info("start index file, path: %s", components.ScanPath)

	err := filepath.Walk(components.ScanPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			// 目录
			processFolder(path)
		} else {
			// 文件
			processFile(path, info)
		}
		return nil
	})
	if err != nil {
		logs.Error("read file fail, path: %s, error: %v", components.ScanPath, err.Error())
		panic(err)
	}
}

// processFolder 处理目录
func processFolder(path string) {
	//fmt.Println(path)
	if strings.Contains(path, "/") {
		split := strings.Split(path, "/")
		folderName := split[len(split)-1]
		fmt.Println(folderName)
		fmt.Println(path)
		// TODO save
	} else {

	}
}

// processFile 处理文件
func processFile(path string, info os.FileInfo) {
	fmt.Println(path, info.Size(), "B", info.Size()/1024, "K", info.Size()/1024/1024, "M")
}
