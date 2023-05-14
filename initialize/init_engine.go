package initialize

import (
	routerGroup "blogger/router"
	"github.com/gin-gonic/gin"
)

func InitEngine() *gin.Engine {
	engine := gin.Default()
	publicGroup := engine.Group("")
	systemRouter := new(routerGroup.RouterGroup).System
	serviceRouter := new(routerGroup.RouterGroup).Service
	{
		systemRouter.InitSystemBase(publicGroup)
		systemRouter.InitSystemDept(publicGroup)
		systemRouter.InitSystemUser(publicGroup)
		serviceRouter.InitFolderRouter(publicGroup)
	}
	return engine
}
