package databaseModels

import "time"

type ForexProductModel struct {
	ExchangeCode      string    `gorm:"column:exchange_code"`
	ProductCode       string    `gorm:"column:product_code"`
	ProductNameTW     string    `gorm:"column:product_name_tw"`
	ProductNameCN     string    `gorm:"column:product_name_cn"`
	ProductNameEN     string    `gorm:"column:product_name_en"`
	ProductPinyin     string    `gorm:"column:product_pinyin"`
	CategoryCode      int       `gorm:"column:category_code"`
	TickUnit          float64   `gorm:"column:tick_unit"`
	MinimumOrder      float64   `gorm:"column:minimum_order"`
	QuoteCurrencyCode string    `gorm:"column:quote_currency_code"`
	ProductStatus     string    `gorm:"column:product_status"`
	Display           string    `gorm:"column:display"`
	IconID            string    `gorm:"column:icon_id"`
	TradeTime         string    `gorm:"column:trade_time"`
	ForexType         int       `gorm:"column:forex_type"` // 0:直接匯兌 1:間接匯兌 2:交叉匯兌(with USD/x) 3:交叉匯兌(with x/USD)
	ForexQuote        string    `gorm:"column:forex_quote"`
	Margin            float64   `gorm:"column:margin"`
	CreatedAt         time.Time `gorm:"column:created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at"`
}
