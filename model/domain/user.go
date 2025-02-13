package domain

type User struct {
	ID         int64      `gorm:"primary_key;column:id"`
	Name       string     `gorm:"column:name"`
	Username   string     `gorm:"column:username"`
	Password   string     `gorm:"column:password"`
	Role       Role       `gorm:"foreignKey:role_id;references:id"`
	BaseDomain BaseDomain `gorm:"embedded"`
}

func (user *User) TableName() string {
	return "m_user"
}
