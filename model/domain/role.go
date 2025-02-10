package domain

type Role struct {
	ID         int64      `gorm:"primary_key;column:id"`
	Name       string     `gorm:"column:name"`
	BaseDomain BaseDomain `gorm:"embedded"`
}

func (role *Role) TableName() string {
	return "m_role"
}
