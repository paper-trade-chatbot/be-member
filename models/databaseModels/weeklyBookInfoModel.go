package databaseModels

import (
	"time"

	"github.com/shopspring/decimal"
)

type WeeklyBookInfoModel struct {
	ID               int64           `gorm:"id"`
	StartDate        time.Time       `gorm:"start_date"`
	EndDate          time.Time       `gorm:"end_date"`
	ABookFloatingPNL decimal.Decimal `gorm:a_book_floating_pnl"`
	ABookHedgePNL    decimal.Decimal `gorm:a_book_hedge_pnl"`
	ABookRealizedPNL decimal.Decimal `gorm:a_book_realized_pnl"`
	ABookTotalPNL    decimal.Decimal `gorm:a_book_total_pnl"`
	ABookNetExposure decimal.Decimal `gorm:a_book_net_exposure"`
	BBookFloatingPNL decimal.Decimal `gorm:b_book_floating_pnl"`
	BBookHedgePNL    decimal.Decimal `gorm:b_book_hedge_pnl"`
	BBookRealizedPNL decimal.Decimal `gorm:b_book_realized_pnl"`
	BBookTotalPNL    decimal.Decimal `gorm:b_book_total_pnl"`
	BBookNetExposure decimal.Decimal `gorm:b_book_net_exposure"`
	CreatedAt        time.Time       `gorm:"column:created_at"`
	UpdatedAt        time.Time       `gorm:"column:updated_at"`
}
