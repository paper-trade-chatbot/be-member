package databaseModels

import (
	"time"

	"github.com/shopspring/decimal"
)

type DailyBookInfoModel struct {
	ID               int64           `gorm:"id"`
	Date             time.Time       `gorm:"date"`
	ABookFloatingPNL decimal.Decimal `gorm:"column:a_book_floating_pnl"`
	ABookHedgePNL    decimal.Decimal `gorm:"column:a_book_hedge_pnl"`
	ABookRealizedPNL decimal.Decimal `gorm:"column:a_book_realized_pnl"`
	ABookTotalPNL    decimal.Decimal `gorm:"column:a_book_total_pnl"`
	ABookNetExposure decimal.Decimal `gorm:"column:a_book_net_exposure"`
	BBookFloatingPNL decimal.Decimal `gorm:"column:b_book_floating_pnl"`
	BBookHedgePNL    decimal.Decimal `gorm:"column:b_book_hedge_pnl"`
	BBookRealizedPNL decimal.Decimal `gorm:"column:b_book_realized_pnl"`
	BBookTotalPNL    decimal.Decimal `gorm:"column:b_book_total_pnl"`
	BBookNetExposure decimal.Decimal `gorm:"column:b_book_net_exposure"`
	CreatedAt        time.Time       `gorm:"column:created_at"`
	UpdatedAt        time.Time       `gorm:"column:updated_at"`
}

type DailyBookInfoSumModel struct {
	ABookFloatingPNL decimal.Decimal `gorm:"column:total_a_book_floating_pnl"`
	ABookHedgePNL    decimal.Decimal `gorm:"column:total_a_book_hedge_pnl"`
	ABookRealizedPNL decimal.Decimal `gorm:"column:total_a_book_realized_pnl"`
	ABookTotalPNL    decimal.Decimal `gorm:"column:total_a_book_total_pnl"`
	ABookNetExposure decimal.Decimal `gorm:"column:total_a_book_net_exposure"`
	BBookFloatingPNL decimal.Decimal `gorm:"column:total_b_book_floating_pnl"`
	BBookHedgePNL    decimal.Decimal `gorm:"column:total_b_book_hedge_pnl"`
	BBookRealizedPNL decimal.Decimal `gorm:"column:total_b_book_realized_pnl"`
	BBookTotalPNL    decimal.Decimal `gorm:"column:total_b_book_total_pnl"`
	BBookNetExposure decimal.Decimal `gorm:"column:total_b_book_net_exposure"`
}

type DailyBookInfoDashboardModel struct {
	ABookTotalPNL    decimal.Decimal `gorm:"column:total_a_book_total_pnl"`
	BBookTotalPNL    decimal.Decimal `gorm:"column:total_b_book_total_pnl"`
	BBookNetExposure decimal.Decimal `gorm:"column:max_b_book_net_exposure"`
}
