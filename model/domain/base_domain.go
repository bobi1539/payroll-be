package domain

import "time"

type BaseDomain struct {
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	CreatedBy     int64     `gorm:"column:created_by"`
	UpdatedBy     int64     `gorm:"column:updated_by"`
	CreatedByName string    `gorm:"column:created_by_name"`
	UpdatedByName string    `gorm:"column:updated_by_name"`
	IsDeleted     bool      `gorm:"column:is_deleted"`
}
