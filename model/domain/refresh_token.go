package domain

import "time"

type RefreshToken struct {
	Id       int64     `gorm:"primary_key;column:id"`
	Token    string    `gorm:"column:token"`
	Validity time.Time `gorm:"column:validity"`
	UserID   int64     `gorm:"column:user_id"`
	User     *User     `gorm:"foreignKey:UserID;references:ID"`
}

func (refreshToken *RefreshToken) TableName() string {
	return "t_refresh_token"
}
