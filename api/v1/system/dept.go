package system

import (
	"blogger/model/system"
	"blogger/service"
	"blogger/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeptApi struct{}

func (d *DeptApi) AddDept(c *gin.Context) {
	var dept system.SysDept
	err := utils.InitApi(c, dept)
	if err != nil {
		return
	}
	logic := service.Service{}.DeptService
	err = logic.AddDept(dept)
	if err != nil {
		utils.FailResCom(c, err, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    1,
	})
}

func (d *DeptApi) GetList(c *gin.Context) {
	var deptForm struct {
		page     int
		pageSize int
		name     string
	}
	err := utils.InitApi(c, deptForm)
	if err != nil {
		return
	}
	utils.InitPage(&deptForm.page, &deptForm.pageSize)
	logic := service.Service{}.DeptService
	list := logic.GetDeptList(deptForm.page, deptForm.pageSize, map[string]interface{}{"name": deptForm.name})
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    list,
	})
}
