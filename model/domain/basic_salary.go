package domain

type BasicSalary struct {
	ID           int64      `gorm:"primary_key;column:id"`
	SalaryAmount int64      `gorm:"column:salary_amount"`
	TotalYear    int32      `gorm:"column:total_year"`
	PositionID   int64      `gorm:"column:position_id"`
	Position     *Position  `gorm:"foreignKey:PositionID;references:ID"`
	BaseDomain   BaseDomain `gorm:"embedded"`
}

func (basicSalary *BasicSalary) TableName() string {
	return "m_basic_salary"
}
