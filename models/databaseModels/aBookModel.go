package databaseModels

import (
	"github.com/shopspring/decimal"
	"time"
)

type ABookModel struct {
	ID                  int64           `gorm:"column:id; primary_key"`
	ExchangeCode        string          `gorm:"column:exchange_code"`
	ProductCode         string          `gorm:"column:product_code"`
	ProductType         int32           `gorm:"column:product_type"`
	BuyPositionQty      decimal.Decimal `gorm:"column:buy_position_qty"`
	SellPositionQty     decimal.Decimal `gorm:"column:sell_position_qty"`
	CurrentPrice        decimal.Decimal `gorm:"column:current_price"`
	Contract            decimal.Decimal `gorm:"column:contract"`
	ExchangeRate        decimal.Decimal `gorm:"column:exchange_rate"`
	EODPrice            decimal.Decimal `gorm:"column:eod_price"`
	HedgePNL            decimal.Decimal `gorm:"column:hedge_pnl"`
	TodayFloatingPNL    decimal.Decimal `gorm:"column:today_floating_pnl"`
	PreviousFloatingPNL decimal.Decimal `gorm:"column:previous_floating_pnl"`
	TodayRealizedPNL    decimal.Decimal `gorm:"column:today_realized_pnl"`
	PreviousRealizedPNL decimal.Decimal `gorm:"column:previous_realized_pnl"`
	CreatedAt           time.Time       `gorm:"column:created_at"`
	UpdatedAt           time.Time       `gorm:"column:updated_at"`
	NetPosition         decimal.Decimal `gorm:"-"`
	GrossPosition       decimal.Decimal `gorm:"-"`
	NetExposure         decimal.Decimal `gorm:"-"`
	TotalPNL            decimal.Decimal `gorm:"-"`
}
