package futurespositiondetaildao

import (
	"github.com/jinzhu/gorm"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
	"github.com/paper-trade-chatbot/be-member/util/specification"
)

const (
	table = "futures_position_detail"
)

//QueryModel lists all queryable columns
type QueryModel struct {
	CustNumber             string
	CustNumberIn           []string
	ExchangeCode           string
	ContractCode           string
	ContractCodeIn         []string
	TradingMemberNumber    string
	TradingMemberNumbertIn []string
	Side                   string
}

func New(tx *gorm.DB, futuresPositionDetail *models.FuturesPositionDetailModel) {
	err := tx.Table(table).
		Create(&futuresPositionDetail).Error

	if err != nil {
		panic(err)
	}
}

func Delete(tx *gorm.DB, order *models.FuturesPositionDetailModel) {
	err := tx.Table(table).
		Where("cust_number = ? AND contract_code = ? AND side = ?", order.CustNumber, order.ContractCode, order.Side).
		Delete(&models.FuturesPositionDetailModel{}).Error

	if err != nil {
		panic(err)
	}
}

func Modify(tx *gorm.DB, order *models.FuturesPositionDetailModel) {
	attrs := map[string]interface{}{
		"exchange_code":                         order.ExchangeCode,
		"trading_member_number":                 order.TradingMemberNumber,
		"open_date":                             order.OpenDate,
		"clearing_date":                         order.ClearingDate,
		"position_qty":                          order.PositionQty,
		"open_price":                            order.OpenPrice,
		"clear_price":                           order.ClearPrice,
		"last_clear_price":                      order.LastClearPrice,
		"position_profit_loss_market_to_market": order.PositionProfitLossMarketToMarket,
		"member_fee":                            order.MemberFee,
		"client_fee":                            order.ClientFee,
		"member_margin":                         order.MemberMargin,
		"client_trade_margin":                   order.ClientTradeMargin,
	}
	err := tx.Table(table).
		Model(models.FuturesPositionDetailModel{}).
		Where("cust_number = ? AND contract_code = ? AND side = ?", order.CustNumber, order.ContractCode, order.Side).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

func GetsWithPagination(tx *gorm.DB, query *QueryModel, paginate *specification.PaginationStruct) ([]*models.FuturesPositionDetailModel, int32) {

	paginate.TableName = table
	var rows []*models.FuturesPositionDetailModel
	var count int32 = 0
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Count(&count).Debug().
		Scopes(specification.NewPaginationSpecification(paginate)).
		Scan(&rows).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil, 0
	}

	if err != nil {
		panic(err)
	}

	return rows, count
}

func Gets(tx *gorm.DB, query *QueryModel) []*models.FuturesPositionDetailModel {
	var result []*models.FuturesPositionDetailModel
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

func Get(tx *gorm.DB, query *QueryModel) *models.FuturesPositionDetailModel {
	result := &models.FuturesPositionDetailModel{}
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

func sideEqualScope(side string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if side != "" {
			return db.Where(table+".side = ?", side)
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
			Scopes(contractCodeEqualScope(query.ContractCode)).
			Scopes(contractCodeIn(query.ContractCodeIn)).
			Scopes(sideEqualScope(query.Side))
	}
}
