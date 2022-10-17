package databaseModels

import "time"

type CryptoProductModel struct {
	ExchangeCode      string    `gorm:"column:exchange_code"`
	ProductCode       string    `gorm:"column:product_code"`
	CryptoName        string    `gorm:"column:crypto_name"`
	ProductNameTW     string    `gorm:"column:product_name_tw"`
	ProductNameCN     string    `gorm:"column:product_name_cn"`
	ProductNameEN     string    `gorm:"column:product_name_en"`
	ProductPinyin     string    `gorm:"column:product_pinyin"`
	CategoryCode      int       `gorm:"column:category_code"`
	TickUnit          float64   `gorm:"column:tick_unit"`
	MinimumOrder      float64   `gorm:"column:minimum_order"`
	IconID            string    `gorm:"column:icon_id"`
	QuoteCurrencyCode string    `gorm:"column:quote_currency_code"`
	ProductStatus     string    `gorm:"column:product_status"`
	Display           string    `gorm:"column:display"`
	TradeTime         string    `gorm:"column:trade_time"`
	CreatedAt         time.Time `gorm:"column:created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at"`
	Description       string    `gorm:"column:description"`
	DescriptionCN     string    `gorm:"column:description_cn"`
	Website           string    `gorm:"column:website"`
}
