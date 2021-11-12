package dao

import (
	"time"

	"github.com/cuilan/file-index/components"
	"github.com/cuilan/file-index/entity"
)

// QueryFolderExist 查询目录是否已存在
func QueryFolderExist(parentId int64, dirName string) (folder entity.Folder) {
	components.Db.Model(&entity.Folder{}).Where("`parent_id` = ? and `name` = ?", parentId, dirName).Find(&folder)
	return folder
}

// AddFolder 添加目录
func AddFolder(parentId int64, dirName, dirPath, dirType, desc string) (id int64) {
	f := &entity.Folder{
		ParentId:   parentId,
		Name:       dirName,
		Path:       dirPath,
		Type:       dirType,
		Removed:    false,
		Hideen:     false,
		Desc:       desc,
		CreateTime: time.Now().Unix() * 1000,
		UpdateTime: time.Now().Unix() * 1000,
	}
	components.Db.Model(&entity.Folder{}).Create(f)
	return f.Id
}
