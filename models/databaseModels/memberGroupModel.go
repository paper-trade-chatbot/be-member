package databaseModels

import (
	"time"
)

type MemberGroupModel struct {
	ID           uint64    `gorm:"column:id; primary_key"`
	Name         string    `gorm:"column:account"`
	Memo         string    `gorm:"column:password_hash"`
	Mail         string    `gorm:"column:mail"`
	LineID       string    `gorm:"column:line_id"`
	CountryCode  string    `gorm:"column:country_code"`
	Phone        string    `gorm:"column:phone"`
	RoleCode     int32     `gorm:"column:role_code"`
	Status       int32     `gorm:"column:status"`
	VerifyStatus int32     `gorm:"column:verify_status"`
	groupID      uint64    `gorm:"column:group_id"`
	CreatedAt    time.Time `gorm:"column:created_at"`
}