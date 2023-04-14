package system

import "time"

type BaseModel struct {
	ID         uint      `gorm:"primaryKey;not null;type:int;autoIncrement" json:"id" `
	CreatedOn  time.Time `json:"created_on" gorm:"comment:创建时间;autoCreateTime"`
	ModifiedOn time.Time `json:"modified_on" gorm:"comment:更新时间; autoUpdateTime"`
}
