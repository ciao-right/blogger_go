package service

import "github.com/gin-gonic/gin"

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

}
