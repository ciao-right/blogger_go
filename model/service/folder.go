package service

import "blogger/model/system"

type Folder struct {
	system.BaseModel
	Name        string   `json:"name,omitempty" validate:"required" gorm:"column:name;type:varchar(20);not null"`
	Level       int      `json:"level,omitempty" gorm:"column:level;type:int;not null"`
	ParentId    int      `json:"parentId,omitempty" gorm:"column:parentId;type:int;not null"`
	Sort        int      `json:"sort,omitempty" gorm:"column:sort;type:int;not null"`
	Icon        string   `json:"icon,omitempty" gorm:"column:icon;type:varchar(255)"` // 图标
	IsHidden    int      `json:"isHidden,omitempty" gorm:"column:isHidden;type:int;not null"`
	Path        string   `json:"path,omitempty" gorm:"column:path;type:varchar(255)"`             // 路径
	Permission  string   `json:"permission,omitempty" gorm:"column:permission;type:varchar(255)"` // 权限标识
	ArticleList []string `json:"article,omitempty" gorm:"column:article;type:varchar(255)"`
	Children    []Folder `json:"children,omitempty" gorm:"-"`
}

func (Folder) TableName() string {
	return "folder"
}
