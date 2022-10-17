package databaseModels

import "time"

type RiskNotifyRecordModel struct {
	ID                uint64    `gorm:"column:id; primary_key"`
	CustNumber        string    `gorm:"column:cust_number"`
	NotifyType        int32     `gorm:"column:notify_type"`
	TotalEquity       float64   `gorm:"column:total_equity"`
	UnrealizedProfit  float64   `gorm:"column:unrealized_profit"`
	WarningLevel      float64   `gorm:"column:warning_level"`
	ReadyToStopLevel  float64   `gorm:"column:ready_to_stop_level"`
	StopLossLevel     float64   `gorm:"column:stop_loss_level"`
	WarningAmount     float64   `gorm:"column:warning_amount"`
	ReadyToStopAmount float64   `gorm:"column:ready_to_stop_amount"`
	StopLossAmount    float64   `gorm:"column:stop_loss_amount"`
	CreatedAt         time.Time `gorm:"column:created_at"`
}
