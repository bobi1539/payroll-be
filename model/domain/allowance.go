package domain

const ALLOWANCE = "Allowance"

type Allowance struct {
	Id         int64      `gorm:"primary_key;column:id"`
	Name       string     `gorm:"column:name"`
	BaseDomain BaseDomain `gorm:"embedded"`
}

func (allowance *Allowance) TableName() string {
	return "m_allowance"
}
