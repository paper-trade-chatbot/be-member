package databaseModels

import (
	"github.com/shopspring/decimal"
	"time"
)

type RiskThroughoutDetailRecordModel struct {
	ID                    uint64          `gorm:"column:id; primary_key"`
	RecordId              string          `gorm:"column:record_id"`
	ExchangeCode          string          `gorm:"column:exchange_code"`
	ProductCode           string          `gorm:"column:product_code"`
	ProductType           int32           `gorm:"column:product_type"`
	AvgPrice              decimal.Decimal `gorm:"column:avg_price"`
	ThroughoutAvgPrice    decimal.Decimal `gorm:"column:throughout_avg_price"`
	ThroughoutPositionQty decimal.Decimal `gorm:"column:throughout_position_qty"`
	ThroughoutLoss        decimal.Decimal `gorm:"column:throughout_loss"`
	Operator              string          `gorm:"column:operator"`
	CreatedAt             time.Time       `gorm:"column:created_at"`
	UpdatedAt             time.Time       `gorm:"column:updated_at"`
}
