package domain

import "time"

type Employee struct {
	Id                int64      `gorm:"primary_key;column:id"`
	Name              string     `gorm:"column:name"`
	PhoneNumber       string     `gorm:"column:phone_number"`
	Email             string     `gorm:"column:email"`
	Address           string     `gorm:"column:address"`
	WorkStatus        string     `gorm:"column:work_status"`
	BankAccountNumber string     `gorm:"column:bank_account_number"`
	BankAccountName   string     `gorm:"column:bank_account_name"`
	Npwp              string     `gorm:"column:npwp"`
	DateOfBirth       time.Time  `gorm:"column:date_of_birth;type:date"`
	JoinDate          time.Time  `gorm:"column:join_date;type:date"`
	IsMarried         bool       `gorm:"column:is_married"`
	TotalChild        int32      `gorm:"column:total_child"`
	PositionId        int64      `gorm:"column:position_id"`
	Position          *Position  `gorm:"foreignKey:PositionId;references:Id"`
	BaseDomain        BaseDomain `gorm:"embedded"`
}

func (employee *Employee) TableName() string {
	return "m_employee"
}
