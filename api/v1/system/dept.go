package system

import (
	"blogger/model/system"
	"blogger/service"
	"blogger/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type DeptApi struct{}

func (d *DeptApi) AddDept(c *gin.Context) {
	var dept system.SysDept
	err := utils.InitPost(c, &dept)
	if err != nil {
		log.Info(err)
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
	}

	var deptSearch system.SysDept
	deptForm.page, _ = strconv.Atoi(c.Query("page"))
	deptForm.pageSize, _ = strconv.Atoi(c.Query("pageSize"))
	deptSearch.Name = c.Query("name")
	utils.InitPage(&deptForm.page, &deptForm.pageSize)
	logic := service.Service{}.DeptService
	list := logic.GetDeptList(deptForm.page, deptForm.pageSize, deptSearch)

	type _deptList struct {
		system.SysDept
		CreatedOn  string `json:"created_on" `
		ModifiedOn string `json:"modified_on" `
	}

	var resDeptList []_deptList
	for _, dept := range list {
		resDeptList = append(resDeptList, _deptList{
			SysDept:    dept,
			CreatedOn:  dept.CreatedOn.Format("2006-01-02 15:04:05"),
			ModifiedOn: dept.ModifiedOn.Format("2006-01-02 15:04:05"),
		})
	}

	data := make(map[string]interface{})
	data["list"] = resDeptList
	data["total"] = logic.GetDeptTotal(deptSearch)
	utils.SuccessResCom(c, data)
}

func (d *DeptApi) UpdateDept(c *gin.Context) {
	var dept system.SysDept
	err := utils.InitPost(c, &dept)
	if err != nil {
		log.Info(err)
		utils.FailResCom(c, err, nil)
		return
	}
	logic := service.Service{}.DeptService
	sErr := logic.UpdateDept(dept)
	if sErr != nil {
		log.Info(sErr)
		utils.FailResCom(c, sErr, nil)
		return
	}
}

func (d *DeptApi) DeleteDept(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	logic := service.Service{}.DeptService
	sErr := logic.DeleteDept(id)
	if sErr != nil {
		log.Info(sErr)
		return
	}
	utils.SuccessResCom(c, 1)
}

func (d *DeptApi) GetTreeList(c *gin.Context) {
	logic := service.Service{}.DeptService
	list := logic.GetDeptTree()
	utils.SuccessResCom(c, list)
}
