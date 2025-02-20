package domain

const POSITION_ALLOWANCE = "PositionAllowance"

type Allowance struct {
	Id              int64          `gorm:"primary_key;column:id"`
	PositionId      int64          `gorm:"column:position_id"`
	Position        *Position      `gorm:"foreignKey:PositionId;references:Id"`
	AllowanceTypeId int64          `gorm:"column:allowance_type_id"`
	AllowanceType   *AllowanceType `gorm:"foreignKey:AllowanceTypeId;references:Id"`
	AllowanceAmount int64          `gorm:"column:allowance_amount"`
	BaseDomain      BaseDomain     `gorm:"embedded"`
}

func (allowance *Allowance) TableName() string {
	return "m_allowance"
}
