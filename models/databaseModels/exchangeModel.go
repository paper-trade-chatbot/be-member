package databaseModels

import "time"

type ExchangeModel struct {
	ID             int64     `gorm:"id"`
	ExchangeCode   string    `gorm:"column:exchange_code"`
	ExchangeName   string    `gorm:"column:exchange_name"`
	ExchangeStatus string    `gorm:"column:exchange_status"`
	TradeTime      string    `gorm:"column:trade_time"`
	ExceptionTime  string    `gorm:"column:exception_time"`
	Timezone       float64   `gorm:"column:timezone"`
	OpenTime       string    `gorm:"column:open_time"`
	Location       string    `gorm:"column:location"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}
