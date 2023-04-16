package system

import "time"

type BaseModel struct {
	ID         uint      `gorm:"primaryKey;not null;type:int;autoIncrement" json:"id" `
	CreatedOn  time.Time `json:"created_on" gorm:"comment:创建时间;autoCreateTime"`
	ModifiedOn time.Time `json:"modified_on" gorm:"comment:更新时间; autoUpdateTime"`
}

type Time interface {
	GetCreatedOn() time.Time
	GetModifiedOn() time.Time
}

func (u *BaseModel) GetCreatedOn() time.Time {
	return u.CreatedOn
}
func (u *BaseModel) GetModifiedOn() time.Time {
	return u.ModifiedOn
}
