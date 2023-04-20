package service

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/cuilan/file-index/configer"
	"github.com/cuilan/file-index/entity"
	"github.com/cuilan/file-index/orm"
	"os"
	"path/filepath"
	"strings"
)

func Index() {
	logs.Info("Application [%s] started...", configer.AppName)
	logs.Info("start index file, path: %s", configer.ScanPath)

	err := filepath.Walk(configer.ScanPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// 目录
			//processFolder(path)
		} else {
			// 文件
			processFile(path, info)
		}
		return nil
	})
	if err != nil {
		logs.Error("read file fail, path: %s, error: %v", configer.ScanPath, err.Error())
		panic(err)
	}
}

// processFolder 处理目录
func processFolder(path string) {
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
	filename := info.Name()
	var suffix = ""
	var hidden = false
	if strings.Contains(filename, ".") {
		if strings.HasPrefix(filename, ".") {
			hidden = true
		}
		split := strings.Split(filename, ".")
		suffix = split[len(split)-1]
	}

	var file entity.File
	orm.SqliteDb.Model(&entity.File{}).Where("hidden = ? and suffix = ? and name = ? and path = ?",
		hidden, suffix, filename, path).First(&file)
	if file.Id != 0 {
		return
	}

	logs.Debug("path: %s", path)
	//fmt.Println(info.Name())
	//fmt.Println(suffix)
	//fmt.Println(info.Size())
	//fmt.Println(info.ModTime().UnixMilli())

	orm.SqliteDb.Create(&entity.File{
		Name:       filename,
		Path:       path,
		Size:       info.Size(),
		Suffix:     suffix,
		Status:     0,
		Removed:    false,
		Hidden:     hidden,
		CreateTime: info.ModTime().UnixMilli(),
		UpdateTime: info.ModTime().UnixMilli(),
	})
}
