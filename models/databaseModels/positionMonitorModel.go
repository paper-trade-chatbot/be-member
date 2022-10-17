package databaseModels

import (
	"github.com/shopspring/decimal"
	"time"
)

type PositionMonitorModel struct {
	ExchangeCode string `gorm:"column:exchange_code"`
	ProductCode  string `gorm:"column:product_code"`
	ProductType  int32  `gorm:"column:product_type"`

	ClientBuyAvgPrice  decimal.Decimal `gorm:"column:client_buy_avg_price"`
	ClientBuyQty       decimal.Decimal `gorm:"column:client_buy_qty"`
	ClientSellAvgPrice decimal.Decimal `gorm:"column:client_sell_avg_price"`
	ClientSellQty      decimal.Decimal `gorm:"column:client_sell_qty"`

	CompanyBuyAvgPrice  decimal.Decimal `gorm:"column:company_buy_avg_price"`
	CompanyBuyQty       decimal.Decimal `gorm:"column:company_buy_qty"`
	CompanySellAvgPrice decimal.Decimal `gorm:"column:company_sell_avg_price"`
	CompanySellQty      decimal.Decimal `gorm:"column:company_sell_qty"`

	TotalBuyAvgPrice  decimal.Decimal `gorm:"column:total_buy_avg_price"`
	TotalSellAvgPrice decimal.Decimal `gorm:"column:total_sell_avg_price"`
	CreatedAt         time.Time       `gorm:"column:created_at"`
	UpdatedAt         time.Time       `gorm:"column:updated_at"`
}
