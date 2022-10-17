package futurespositiondao

import (
	"github.com/jinzhu/gorm"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
)

const (
	table = "futures_position"
)

//QueryModel lists all queryable columns
type QueryModel struct {
	CustNumber             string
	CustNumberIn           []string
	ContractCode           string
	ContractCodeIn         []string
	TradingMemberNumber    string
	TradingMemberNumbertIn []string
	TradeType              string
	ExchangeCode           string
}

func New(tx *gorm.DB, futuresPosition *models.FuturesPositionModel) {
	err := tx.Table(table).
		Create(&futuresPosition).Error

	if err != nil {
		panic(err)
	}
}

func Delete(tx *gorm.DB, futuresPosition *models.FuturesPositionModel) {
	err := tx.Table(table).Where("cust_number=? AND contract_code=?", futuresPosition.CustNumber, futuresPosition.ContractCode).Delete(&models.FuturesPositionModel{}).Error

	if err != nil {
		panic(err)
	}
}

func Modify(tx *gorm.DB, order *models.FuturesPositionModel) {
	attrs := map[string]interface{}{
		"clearing_date":                         order.ClearingDate,
		"exchange_code":                         order.ExchangeCode,
		"trading_member_number":                 order.TradingMemberNumber,
		"trade_type":                            order.TradeType,
		"clear_price":                           order.ClearPrice,
		"last_clear_price":                      order.LastClearPrice,
		"buy_open_qty":                          order.BuyOpenQty,
		"buy_close_qty":                         order.BuyCloseQty,
		"buy_trade_qty":                         order.BuyTradeQty,
		"sell_open_qty":                         order.SellOpenQty,
		"sell_close_qty":                        order.SellCloseQty,
		"sell_trade_qty":                        order.SellTradeQty,
		"position_profit_loss_market_to_market": order.PositionProfitLossMarketToMarket,
		"member_fee":                            order.MemberFee,
		"client_fee":                            order.ClientFee,
		"premium_in":                            order.PremiumIn,
		"premium_out":                           order.PremiumOut,
		"member_margin":                         order.MemberMargin,
		"client_trade_margin":                   order.ClientTradeMargin,
	}
	err := tx.Table(table).
		Model(models.FuturesPositionModel{}).
		Where("cust_number = ? AND contract_code = ?", order.CustNumber, order.ContractCode).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

func Gets(tx *gorm.DB, query *QueryModel) []*models.FuturesPositionModel {
	var result []*models.FuturesPositionModel
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}

	if err != nil {
		panic(err)
	}

	return result
}

func Get(tx *gorm.DB, query *QueryModel) *models.FuturesPositionModel {
	result := &models.FuturesPositionModel{}
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scan(result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}

	if err != nil {
		panic(err)
	}
	return result
}

func custNumberEqualScope(account string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if account != "" {
			return db.Where(table+".cust_number = ?", account)
		}
		return db
	}
}

func custNumberInScope(custNumber []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(custNumber) > 0 {
			return db.Where(table+".cust_number IN (?)", custNumber)
		}
		return db
	}
}

func tradingMemberNumEqualScope(tradingMemberNum string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if tradingMemberNum != "" {
			return db.Where(table+".trading_member_number = ?", tradingMemberNum)
		}
		return db
	}
}

func tradingMemberNumInScope(tradingMemberNumIn []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(tradingMemberNumIn) > 0 {
			return db.Where(table+".trading_member_number IN (?)", tradingMemberNumIn)
		}
		return db
	}
}

func contractCodeEqualScope(contractCode string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if contractCode != "" {
			return db.Where(table+".contract_code = ?", contractCode)
		}
		return db
	}
}

func exchangeCodeEqualScope(exchangeCode string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if exchangeCode != "" {
			return db.Where(table+".exchange_code = ?", exchangeCode)
		}
		return db
	}
}

func tradeTypeEqualScope(tradeType string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if tradeType != "" {
			return db.Where(table+".trade_type = ?", tradeType)
		}
		return db
	}
}

func contractCodeIn(contractCode []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(contractCode) > 0 {
			return db.Where(table+".contract_code IN (?)", contractCode)
		}
		return db
	}
}

func queryChain(query *QueryModel) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Scopes(custNumberEqualScope(query.CustNumber)).
			Scopes(custNumberInScope(query.CustNumberIn)).
			Scopes(tradingMemberNumEqualScope(query.TradingMemberNumber)).
			Scopes(tradingMemberNumInScope(query.TradingMemberNumbertIn)).
			Scopes(exchangeCodeEqualScope(query.ExchangeCode)).
			Scopes(tradeTypeEqualScope(query.TradeType)).
			Scopes(contractCodeEqualScope(query.ContractCode)).
			Scopes(contractCodeIn(query.ContractCodeIn))
	}
}
