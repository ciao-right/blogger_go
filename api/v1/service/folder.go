package service

import (
	"blogger/model/service"
	"blogger/model/system"
	svc "blogger/service/service"
	"blogger/utils"
	"github.com/gin-gonic/gin"
)

type FolderApi struct{}

// @Summary 新增文件夹
// @Description 新增文件夹
// @Tags 文件夹
// @Accept  json
// @Produce  json
// @Param   name     query    string     true        "文件夹名称"
// @Param   parentId     query    int     true        "父级ID"
// @Param   sort     query    int     true        "排序"
// @Param   icon     query    string     true        "图标"
// @Param   path     query    string     true        "路径"
// @Param   isHidden     query    int     true        "是否隐藏"
// @Param   permission     query    string     true        "权限标识"

// AddFolder @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/folder/add [post]
func (FolderApi) AddFolder(c *gin.Context) {
	var folder service.Folder
	err := utils.InitPost(c, &folder)
	if err != nil {
		utils.FailResCom(c, err, nil)
		return
	}
	logic := svc.FolderService{}
	err = logic.AddFolder(folder)
	if err != nil {
		utils.FailResCom(c, err, nil)
		return
	}

	utils.SuccessResCom(c, 1)
}

func (FolderApi) GetFolderList(c *gin.Context) {
	var folderSearch system.SearchPage
	err := utils.InitPost(c, &folderSearch)
	if err != nil {
		utils.FailResCom(c, err, nil)
		return
	}
	logic := svc.FolderService{}
	folderList, err := logic.GetFolderList(folderSearch)
	if err != nil {
		utils.FailResCom(c, err, nil)
		return
	}
	utils.SuccessResCom(c, folderList)

}

func (FolderApi) UpdateFolder(c *gin.Context) {
	var folder service.Folder
	err := utils.InitPost(c, &folder)
	if err != nil {
		utils.FailResCom(c, err, nil)
		return
	}
	logic := svc.FolderService{}
	err = logic.UpdateFolder(folder)
	if err != nil {
		utils.FailResCom(c, err, nil)
		return
	}
	utils.SuccessResCom(c, 1)
}

func (FolderApi) DeleteFolder(c *gin.Context) {
	id := c.Param("id")
	logic := svc.FolderService{}
	err := logic.DeleteFolder(id)
	if err != nil {
		utils.FailResCom(c, err, nil)
		return
	}
	utils.SuccessResCom(c, 1)

}
