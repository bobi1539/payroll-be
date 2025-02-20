package domain

const ALLOWANCE_TYPE = "AllowanceType"

type AllowanceType struct {
	Id         int64      `gorm:"primary_key;column:id"`
	Name       string     `gorm:"column:name"`
	BaseDomain BaseDomain `gorm:"embedded"`
}

func (allowanceType *AllowanceType) TableName() string {
	return "m_allowance_type"
}
