package databaseModels

import "time"

type FuturesPositionModel struct {
	CustNumber                       string    `gorm:"column:cust_number"`
	ContractCode                     string    `gorm:"column:contract_code"`
	ClearingDate                     string    `gorm:"column:clearing_date"`
	ExchangeCode                     string    `gorm:"column:exchange_code"`
	TradingMemberNumber              string    `gorm:"column:trading_member_number"`
	TradeType                        string    `gorm:"column:trade_type"`
	ClearPrice                       float64   `gorm:"column:clear_price"`
	LastClearPrice                   float64   `gorm:"column:last_clear_price"`
	BuyOpenQty                       float64   `gorm:"column:buy_open_qty"`
	BuyCloseQty                      float64   `gorm:"column:buy_close_qty"`
	BuyTradeQty                      float64   `gorm:"column:buy_trade_qty"`
	SellOpenQty                      float64   `gorm:"column:sell_open_qty"`
	SellCloseQty                     float64   `gorm:"column:sell_close_qty"`
	SellTradeQty                     float64   `gorm:"column:sell_trade_qty"`
	BuyTradePrice                    float64   `gorm:"column:buy_trade_price"`
	SellTradePrice                   float64   `gorm:"column:sell_trade_price"`
	PositionProfitLossMarketToMarket float64   `gorm:"column:position_profit_loss_market_to_market"`
	MemberFee                        float64   `gorm:"column:member_fee"`
	ClientFee                        float64   `gorm:"column:client_fee"`
	PremiumIn                        float64   `gorm:"column:premium_in"`
	PremiumOut                       float64   `gorm:"column:premium_out"`
	MemberMargin                     float64   `gorm:"column:member_margin"`
	ClientTradeMargin                float64   `gorm:"column:client_trade_margin"`
	CreatedAt                        time.Time `gorm:"column:created_at"`
	UpdatedAt                        time.Time `gorm:"column:updated_at"`
}
