package databaseModels

import "time"

type CronjobModel struct {
	ID        int64     `gorm:"id"`
	Log       string    `gorm:"column:log"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
