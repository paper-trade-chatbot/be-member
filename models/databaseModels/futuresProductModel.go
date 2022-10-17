package databaseModels

import "time"

type FuturesProductModel struct {
	ExchangeCode         string  `gorm:"column:exchange_code"`
	ProductCode          string  `gorm:"column:product_code"`
	ProductName          string  `gorm:"column:product_name"`
	ProductPinyin        string  `gorm:"column:product_pinyin"`
	ProductStatus        string  `gorm:"column:product_status"`
	Display              string  `gorm:"column:display"`
	FuturesVarietyCode   string  `gorm:"column:futures_variety_code"`
	FuturesVarietyCodeID int32   `gorm:"column:futures_variety_code_id"`
	OptionSetCode        string  `gorm:"column:option_set_code"`
	QuoteCurrencyCode    string  `gorm:"column:quote_currency_code"`
	ClearCurrencyCode    string  `gorm:"column:clear_currency_code"`
	InstrumentMultiple   float64 `gorm:"column:instrument_multiple"`
	FaceValue            string  `gorm:"column:face_value"`
	PriceUnit            float64 `gorm:"column:price_unit"`
	TickUnit             float64 `gorm:"column:tick_unit"`
	IconID               string  `gorm:"column:icon_id"`
	MinimumOrder         float64 `gorm:"column:minimum_order"`

	LastTradeDate    int       `gorm:"column:last_trade_date"`
	UnderlyingID     float64   `gorm:"column:underlying_id"`
	UnderlyingUnit   float64   `gorm:"column:underlying_unit"`
	OptionType       float64   `gorm:"column:option_type"`
	ExecuteStartDate int       `gorm:"column:execute_start_date"`
	ExecuteLastDate  int       `gorm:"column:execute_last_date"`
	LastClearPrice   float64   `gorm:"column:last_clear_price"`
	ClearPrice       float64   `gorm:"column:clear_price"`
	MinHoldBalance   float64   `gorm:"column:min_hold_balance"`
	InstrumentType   int       `gorm:"column:instrument_type"`
	TradeTime        string    `gorm:"column:trade_time"`
	Margin           float64   `gorm:"column:margin"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}
