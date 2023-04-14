package system

type User struct {
	BaseModel `gorm:"embedded"`
	Account   string  `json:"account" gorm:"column:account; type:varchar(30);not null; "`
	Password  string  `json:"password" gorm:"column:password; type:varchar(30);not null"`
	Name      string  `json:"name" gorm:"column:name;type:varchar(20) ;not null"`
	Age       int     `json:"age" gorm:"column:age; type:int(4)"`
	Sex       int     `json:"sex" gorm:"column:sex;type:int(2)" `
	Phone     string  `json:"phone" gorm:"column:phone;not null;type:int(13)" `
	DeptId    SysDept `json:"deptId" gorm:"column:deptId;not null;type:int"`
	Remark    string  `json:"remark" gorm:"column:remark;type:varchar(50)"`
}

func (u *User) TableName() string {
	return "sys_user"
}
