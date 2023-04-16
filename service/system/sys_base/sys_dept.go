package sys_base

import (
	"blogger/global"
	"blogger/model/system"
	"errors"
	"fmt"
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

func (d *DeptService) GetDeptList(page, pageSize int, condition interface{}) (deptList []system.SysDept) {
	db := global.Db
	fmt.Println(condition)
	db.Model(&system.SysDept{}).Where(condition).Limit(pageSize).Offset((page - 1) * pageSize).Order("sort").Find(&deptList)
	return
}

func (d *DeptService) GetDeptTotal(condition interface{}) (total int64) {
	db := global.Db
	db.Model(&system.SysDept{}).Where(condition).Count(&total)
	return total
}

func (d *DeptService) UpdateDept(dept system.SysDept) error {
	db := global.Db
	if dept.ID == 0 {
		return errors.New("id不能为空")
	}
	res := db.Model(&system.SysDept{}).Where("id = ?", dept.ID).Updates(dept)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (d *DeptService) DeleteDept(id int) error {
	db := global.Db
	if id == 0 {
		return errors.New("id不能为空")
	}
	res := db.Where("id = ?", id).Delete(&system.SysDept{})
	return res.Error
}

func hasSameDeptName(name string) bool {
	db := global.Db
	var d system.SysDept
	res := db.Model(&system.SysDept{}).Where("name = ?", name).First(&d)
	return res.RowsAffected > 0
}

func (d *DeptService) GetDeptTree() (deptTree []system.SysDept) {
	db := global.Db
	db.Model(&system.SysDept{}).Where("parent = ?", 0).Not("state = 0").Order("sort").Find(&deptTree)
	for i := 0; i < len(deptTree); i++ {
		deptTree[i].Children = getDeptChildren(deptTree[i].ID)
	}
	return
}

func getDeptChildren(parentId uint) (deptChildren []system.SysDept) {
	db := global.Db
	db.Model(&system.SysDept{}).Where("parent = ?", parentId).Not("state = 0").Order("sort").Find(&deptChildren)
	for i := 0; i < len(deptChildren); i++ {
		deptChildren[i].Children = getDeptChildren(deptChildren[i].ID)
	}
	return
}
