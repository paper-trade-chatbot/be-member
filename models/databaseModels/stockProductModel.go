package databaseModels

import "time"

type StockProductModel struct {
	ExchangeCode      string    `gorm:"column:exchange_code"`
	ProductCode       string    `gorm:"column:product_code"`
	ProductNameEN     string    `gorm:"column:product_name_en"`
	ProductPinyin     string    `gorm:"column:product_pinyin"`
	ProductStatus     string    `gorm:"column:product_status"`
	Display           string    `gorm:"column:display"`
	QuoteCurrencyCode string    `gorm:"column:quote_currency_code"`
	TickUnit          float64   `gorm:"column:tick_unit"`
	MinimumOrder      float64   `gorm:"column:minimum_order"`
	IconID            string    `gorm:"column:icon_id"`
	CreatedAt         time.Time `gorm:"column:created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at"`
}
