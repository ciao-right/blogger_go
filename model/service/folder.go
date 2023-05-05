package service

import "blogger/model/system"

type Folder struct {
	system.BaseModel
	Name     string   `json:"name,omitempty"`
	Level    int      `json:"level,omitempty"`
	Children []Folder `json:"children,omitempty"`
}

func (Folder) TableName() string {
	return "folder"
}
