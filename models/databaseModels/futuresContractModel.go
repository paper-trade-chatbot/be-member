package databaseModels

import "time"

type FuturesContractModel struct {
	ExchangeCode       string    `gorm:"column:exchange_code"`
	ContractCode       string    `gorm:"column:contract_code"`
	ContractStatus     string    `gorm:"column:contract_status"`
	FutureVarietyCode  string    `gorm:"column:future_variety_code"`
	OptionSetCode      string    `gorm:"column:option_set_code"`
	QuoteCurrencyCode  string    `gorm:"column:quote_currency_code"`
	ClearCurrencyCode  string    `gorm:"column:clear_currency_code"`
	InstrumentMultiple float64   `gorm:"column:instrument_multiple"`
	FaceValue          string    `gorm:"column:face_value"`
	PriceUnit          int       `gorm:"column:price_unit"`
	TickUnit           int       `gorm:"column:tick_unit"`
	Month              int       `gorm:"column:month"`
	LastTradeDate      int       `gorm:"column:last_trade_date"`
	UnderlyingID       float64   `gorm:"column:underlying_id"`
	UnderlyingUnit     float64   `gorm:"column:underlying_unit"`
	OptionType         float64   `gorm:"column:option_type"`
	ExecuteStartDate   int       `gorm:"column:execute_start_date"`
	ExecuteLastDate    int       `gorm:"column:execute_last_date"`
	LastClearPrice     float64   `gorm:"column:last_clear_price"`
	ClearPrice         float64   `gorm:"column:clear_price"`
	MinHoldBalance     float64   `gorm:"column:min_hold_balance"`
	InstrumentType     int       `gorm:"column:instrument_type"`
	CreatedAt          time.Time `gorm:"column:created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at"`
}
