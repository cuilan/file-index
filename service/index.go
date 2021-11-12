package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/cuilan/file-index/components"
	"github.com/cuilan/file-index/dao"
)

var (
	Separator string = "/"
)

func Index() {
	logs.Info("Application [%s] started...", components.AppName)
	logs.Info("start index file, path: %s", components.ScanPath)

	// 配置中的目录包含 "/" 说明不是根目录，需要特殊处理
	parentId := baseDir(components.ScanPath)
	logs.Info("config scan_path: %s, parentId: %d", components.ScanPath, parentId)

	// 处理配置目录
	readDir(parentId, components.ScanPath)
}

// baseDir 处理基础目录
func baseDir(dir string) (parentId int64) {
	if strings.Contains(dir, Separator) {
		logs.Info("app.conf scan_path is not root path, start process...")
		parentId = 0
		var dirName, dirPath, dirType, desc string
		dirs := strings.Split(dir, Separator)
		for i := range dirs {
			dirName = dirs[i]
			existDir := dao.QueryFolderExist(parentId, dirName)
			// id=0，则目录不存在
			if existDir.Id != int64(0) {
				parentId = existDir.Id
				continue
			}
			// 根目录
			if parentId == 0 {
				dirPath = dirName
				dirType = "盘符"
				desc = "根目录"

				logs.Info("baseDir create root dir: %s", dirName)
				parentId = dao.AddFolder(parentId, dirName, dirPath, dirType, desc)
				dirType = ""
				desc = ""
			} else {
				// 如果当前目录名称不是根目录，则需要拼接绝对路径
				dirPath = path.Join(dirPath, dirName)
				logs.Info("baseDir create dir: %s", dirName)
				parentId = dao.AddFolder(parentId, dirName, dirPath, dirType, desc)
			}
		}
	}
	return
}

func readDir(parentId int64, dir string) {
	logs.Info("read dir: %s, parentId: %d", dir, parentId)
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		logs.Error("read dir: %s, error: %v", dir, err.Error())
		return
	}
	for i := range infos {
		info := infos[i]
		// 处理目录
		if info.IsDir() {
			var dirName, dirPath, dirType, desc string
			dirName = info.Name()

			// 忽略隐藏目录
			if dirName[0:1] == "." {
				continue
			}

			// 拼接绝对路径
			dirPath = path.Join(dir, dirName)
			existDir := dao.QueryFolderExist(parentId, dirName)
			var currentId int64
			if existDir.Id == 0 {
				// id=0，则目录不存在
				logs.Info("readDir create dir: %s", dirName)
				currentId = dao.AddFolder(parentId, dirName, dirPath, dirType, desc)
			} else {
				// id!=0，则目录已存在
				currentId = existDir.Id
			}

			readDir(currentId, dirPath)
		} else {
			// TODO 目录已存在，处理文件
		}
	}
}

// processFile 处理文件
func processFile(path string, info os.FileInfo) {
	fmt.Println(path, info.Size(), "B", info.Size()/1024, "K", info.Size()/1024/1024, "M")
}
