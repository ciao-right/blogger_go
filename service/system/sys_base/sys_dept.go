package sys_base

import (
	"blogger/global"
	"blogger/model/system"
	"errors"
)

type DeptService struct{}

func (d *DeptService) AddDept(dept system.SysDept) error {
	db := global.Db
	if hasSameDeptName(dept.Name) {
		return errors.New("已经存在相同名称的部门")
	} else {
		res := db.Create(&dept)
		if res.Error != nil {
			return res.Error
		} else {
			return nil
		}
	}
}
func hasSameDeptName(name string) bool {
	db := global.Db
	var d system.SysDept
	res := db.Model(&system.SysDept{}).Where("name = ?", name).First(&d)
	return res.RowsAffected > 0
}
func (d *DeptService) GetDeptList(page, pageSize int, condition map[string]interface{}) (deptList []system.SysDept) {
	db := global.Db
	db.Model(&system.SysDept{}).Where(condition).Limit(pageSize).Offset((page - 1) * pageSize).Find(&deptList)
	return
}
