package serviceGroup

import (
	v1 "blogger/api/v1"
	"github.com/gin-gonic/gin"
)

type FolderRouter struct{}

func (FolderRouter) InitFolderRouter(group *gin.RouterGroup) gin.IRouter {
	folderRouter := group.Group("folder")
	api := new(v1.Api).Service.FolderApi
	{
		folderRouter.POST("/addFolder", api.AddFolder)
		folderRouter.POST("/getFolderList", api.GetFolderList)
		folderRouter.POST("/updateFolder", api.UpdateFolder)
		folderRouter.GET("/deleteFolder", api.DeleteFolder)

	}
	return folderRouter
}
