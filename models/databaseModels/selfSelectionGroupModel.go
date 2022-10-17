package databaseModels

import "time"

type SelfSelectionGroupModel struct {
	CustNumber string    `gorm:"column:cust_number"`
	Group      string    `gorm:"column:group"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}
