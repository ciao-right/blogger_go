package system

type SysDept struct {
	BaseModel `gorm:"embedded"`
	Name      string    `json:"name" gorm:"column:name;not null;type:varchar(15)"`
	State     int       `json:"state" gorm:"column:state;not null; type:int;default:1"`
	Parent    int       `json:"parent" gorm:"column:parent;not null;type:int;default:0"`
	Sort      int       `json:"sort" gorm:"column:sort; not null ;type:int;default:1"`
	Children  []SysDept `json:"children" gorm:"-"`
}

func (u *SysDept) TableName() string {
	return "sys_dept"
}
