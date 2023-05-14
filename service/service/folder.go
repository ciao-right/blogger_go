package service

import (
	"blogger/global"
	"blogger/model/service"
	"blogger/model/system"
	"errors"
)

type FolderService struct{}

func (FolderService) AddFolder(folder service.Folder) error {
	db := global.Db
	if hasSameFolder(folder.Name) {
		return errors.New("已经存在相同的文件名")
	}
	result := db.Create(&folder)
	return result.Error
}

func hasSameFolder(name string) bool {
	db := global.Db
	var folder service.Folder
	db.Model(service.Folder{}).Where("name = ?", name).First(&folder)
	return folder.Name != ""
}

func (FolderService) GetFolderList(search system.SearchPage) ([]service.Folder, error) {
	db := global.Db
	var folderList []service.Folder
	result := db.Model(service.Folder{}).Limit(search.PageSize).Offset((search.Page - 1) * search.PageSize).Find(&folderList)
	return folderList, result.Error
}
