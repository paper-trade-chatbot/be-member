package databaseModels

import (
	"time"

	"github.com/shopspring/decimal"
)

type PositionDetailModel struct {
	CustNumber   string `gorm:"primary_key;column:cust_number"`
	Side         int32  `gorm:"primary_key;column:side"`
	ExchangeCode string `gorm:"column:exchange_code"`
	ProductCode  string `gorm:"primary_key;column:product_code"`
	ProductType  int32  `gorm:"primary_key;column:product_type"`

	Leverage       int             `gorm:"column:leverage"`
	SelfPaidMargin decimal.Decimal `gorm:"column:self_paid_margin"`
	LeverageMargin decimal.Decimal `gorm:"column:leverage_margin"`

	PositionQty     decimal.Decimal `gorm:"column:position_qty"`
	OpenPrice       decimal.Decimal `gorm:"column:open_price"`
	ClearPrice      decimal.Decimal `gorm:"column:clear_price"`
	AveragePrice    decimal.Decimal `gorm:"column:average_price"`
	LastClearPrice  decimal.Decimal `gorm:"column:last_clear_price"`
	StopLossPrice   decimal.Decimal `gorm:"column:stop_loss_price"`
	GainProfitPrice decimal.Decimal `gorm:"column:gain_profit_price"`

	ClosePriceAndMargin string    `gorm:"column:close_price_and_margin"`
	CreatedAt           time.Time `gorm:"column:created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at"`
}
