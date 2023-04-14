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
	}

	return gin
}
