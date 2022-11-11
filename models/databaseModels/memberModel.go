package databaseModels

import (
	"time"

	"gorm.io/gorm"
)

type MemberModel struct {
	ID           uint64         `gorm:"column:id; primary_key"`
	Account      string         `gorm:"column:account"`
	PasswordHash string         `gorm:"column:password_hash"`
	Mail         string         `gorm:"column:mail"`
	LineID       string         `gorm:"column:line_id"`
	CountryCode  string         `gorm:"column:country_code"`
	Phone        string         `gorm:"column:phone"`
	Status       int32          `gorm:"column:status"`
	VerifyStatus int32          `gorm:"column:verify_status"`
	GroupID      uint64         `gorm:"column:group_id"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`
}
