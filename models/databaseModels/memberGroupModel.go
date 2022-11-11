package databaseModels

import (
	"time"
)

type MemberGroupModel struct {
	ID        uint64    `gorm:"column:id; primary_key"`
	Name      string    `gorm:"column:name"`
	Memo      string    `gorm:"column:memo"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
