package system

type User struct {
	BaseModel `gorm:"embedded"`
	Account   string  `json:"account" gorm:"column:account; type:varchar(30);not null;" validate:"required" `
	Password  string  `json:"password" gorm:"column:password; type:varchar(255);not null" validate:"required"`
	Name      string  `json:"name" gorm:"column:name;type:varchar(20) ;not null" validate:"required"`
	Age       int     `json:"age" gorm:"column:age; type:int(4)"`
	Sex       int     `json:"sex" gorm:"column:sex;type:int(2)" validate:"oneof= 0 1"`
	Phone     string  `json:"phone" gorm:"column:phone;not null;type:varchar(13)" validate:"required"`
	DeptID    int     `json:"deptId" gorm:"column:deptId;not null;type:int" validate:"required"`
	Remark    string  `json:"remark" gorm:"column:remark;type:varchar(50)"`
	Dept      SysDept `json:"dept" gorm:"foreignKey:DeptID"`
}

func (u *User) TableName() string {
	return "sys_user"
}
