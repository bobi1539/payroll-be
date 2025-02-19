package domain

const POSITION = "Position"

type Position struct {
	ID         int64      `gorm:"primary_key;column:id"`
	Name       string     `gorm:"column:name"`
	BaseDomain BaseDomain `gorm:"embedded"`
}

func (position *Position) TableName() string {
	return "m_position"
}
