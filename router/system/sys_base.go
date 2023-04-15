package systemGroup

import (
	v1 "blogger/api/v1"
	"github.com/gin-gonic/gin"
)

type sysBaseRouter struct{}

func (s *sysBaseRouter) InitSystemBase(group *gin.RouterGroup) gin.IRouter {
	baseRouter := group.Group("base") //都是以/base 为开头
	baseApi := new(v1.Api).System.BaseApi
	{
		baseRouter.POST("/login", baseApi.Login)
		baseRouter.POST("/register", baseApi.Register)
	}
	return baseRouter
}
