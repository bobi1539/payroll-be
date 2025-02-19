package domain

const USER = "User"

type User struct {
	Id         int64      `gorm:"primary_key;column:id"`
	Name       string     `gorm:"column:name"`
	Username   string     `gorm:"column:username"`
	Password   string     `gorm:"column:password"`
	RoleId     int64      `gorm:"column:role_id"`
	Role       *Role      `gorm:"foreignKey:RoleId;references:Id"`
	BaseDomain BaseDomain `gorm:"embedded"`
}

func (user *User) TableName() string {
	return "m_user"
}
