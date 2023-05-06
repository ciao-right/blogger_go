package service

import "blogger/model/system"

type Article struct {
	system.BaseModel
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	// 0:草稿 1:已发布
	State int `json:"state,omitempty"`
	// 0:原创 1:转载
	Type int `json:"type,omitempty"`
	// 0:公开 1:私密
	Permission int `json:"permission,omitempty"`
	// 0:不置顶 1:置顶
	Top int `json:"top,omitempty"`
}
