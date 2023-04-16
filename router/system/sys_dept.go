package systemGroup

import (
	v1 "blogger/api/v1"
	"github.com/gin-gonic/gin"
)

type sysDeptRouter struct{}

func (s *sysBaseRouter) InitSystemDept(group *gin.RouterGroup) gin.IRouter {
	baseRouter := group.Group("base") //都是以/base 为开头
	deptApi := new(v1.Api).System.DeptApi
	{
		baseRouter.POST("/addDept", deptApi.AddDept)
		baseRouter.GET("/getBaseDept", deptApi.GetList)
		baseRouter.POST("/updateDept", deptApi.UpdateDept)
		baseRouter.GET("/deleteDept", deptApi.DeleteDept)
		baseRouter.GET("/getDeptTree", deptApi.GetTreeList)
	}
	return baseRouter
}
