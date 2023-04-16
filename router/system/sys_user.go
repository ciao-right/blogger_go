package systemGroup

import (
	v1 "blogger/api/v1"
	"github.com/gin-gonic/gin"
)

type sysUserRouter struct{}

func (s *sysBaseRouter) InitSystemUser(group *gin.RouterGroup) gin.IRouter {
	baseRouter := group.Group("base") //都是以/base 为开头
	deptApi := new(v1.Api).System.UserApi
	{
		baseRouter.POST("/getUser", deptApi.GetList)
		baseRouter.POST("/updateUser", deptApi.Update)
		baseRouter.GET("/deleteUser", deptApi.Delete)
	}
	return baseRouter
}
