package databaseModels

import "time"

type FuturesPositionDetailModel struct {
	CustNumber                       string     `gorm:"column:cust_number"`
	ContractCode                     string     `gorm:"column:contract_code"`
	Side                             int        `gorm:"column:side"`
	ExchangeCode                     string     `gorm:"column:exchange_code"`
	TradingMemberNumber              string     `gorm:"column:trading_member_number"`
	OpenDate                         *time.Time `gorm:"column:open_date"`
	ClearingDate                     *time.Time `gorm:"column:clearing_date"`
	PositionQty                      float64    `gorm:"column:position_qty"`
	OpenPrice                        float64    `gorm:"column:open_price"`
	ClearPrice                       float64    `gorm:"column:clear_price"`
	LastClearPrice                   float64    `gorm:"column:last_clear_price"`
	PositionProfitLossMarketToMarket float64    `gorm:"column:position_profit_loss_market_to_market"`
	MemberFee                        float64    `gorm:"column:member_fee"`
	ClientFee                        float64    `gorm:"column:client_fee"`
	MemberMargin                     float64    `gorm:"column:member_margin"`
	ClientTradeMargin                float64    `gorm:"column:client_trade_margin"`
	CreatedAt                        time.Time  `gorm:"column:created_at"`
	UpdatedAt                        time.Time  `gorm:"column:updated_at"`
}
