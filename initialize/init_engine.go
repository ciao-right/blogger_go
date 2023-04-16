package initialize

import (
	routerGroup "blogger/router"
	"github.com/gin-gonic/gin"
)

func InitEngine() *gin.Engine {
	gin := gin.Default()
	publicGroup := gin.Group("")
	systemRouter := new(routerGroup.RouterGroup).System
	{
		systemRouter.InitSystemBase(publicGroup)
		systemRouter.InitSystemDept(publicGroup)
		systemRouter.InitSystemUser(publicGroup)
	}

	return gin
}
