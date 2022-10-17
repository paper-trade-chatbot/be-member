package positiondetaildao

import (
	"github.com/jinzhu/gorm"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
)

const (
	table = "position_detail"
)

//QueryModel lists all queryable columns
type QueryModel struct {
	CustNumber             string
	CustNumberIn           []string
	ExchangeCode           string
	ProductCode            string
	ProductCodeIn          []string
	TradingMemberNumber    string
	TradingMemberNumbertIn []string
	Side                   string
	ProductType            int
	ProductTypeIn          []int
	ForUpdate              bool
}

func New(tx *gorm.DB, stockPositionDetail *models.PositionDetailModel) {
	err := tx.Table(table).
		Create(&stockPositionDetail).Error

	if err != nil {
		panic(err)
	}
}

func Delete(tx *gorm.DB, order *models.PositionDetailModel) {
	err := tx.Table(table).
		Where("cust_number = ? AND product_code = ? AND side = ?", order.CustNumber, order.ProductCode, order.Side).
		Delete(&models.PositionDetailModel{}).Error

	if err != nil {
		panic(err)
	}
}

func Gets(tx *gorm.DB, query *QueryModel) []*models.PositionDetailModel {
	var result []*models.PositionDetailModel
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

func Get(tx *gorm.DB, query *QueryModel) *models.PositionDetailModel {
	result := &models.PositionDetailModel{}
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

func GetsDistinctProductCode(tx *gorm.DB, query *QueryModel) []*models.PositionDetailModel {
	var result []*models.PositionDetailModel
	err := tx.Table(table).
		Select("distinct product_code").Debug().
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

func productCodeEqualScope(productCode string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if productCode != "" {
			return db.Where(table+".product_code = ?", productCode)
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

func productCodeIn(productCode []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(productCode) > 0 {
			return db.Where(table+".product_code IN (?)", productCode)
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

func forUpdateScope(forUpdate bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if forUpdate {
			return db.Set("gorm:query_option", "FOR UPDATE")
		}
		return db
	}
}

func productTypeInScope(productTypeIn []int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(productTypeIn) > 0 {
			return db.Where(table+".product_type IN (?)", productTypeIn)
		}
		return db
	}
}

func productTypeEqualScope(productType int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if productType != 0 {
			return db.Where(table+".product_type = ?", productType)
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
			Scopes(productCodeEqualScope(query.ProductCode)).
			Scopes(productCodeIn(query.ProductCodeIn)).
			Scopes(sideEqualScope(query.Side)).
			Scopes(forUpdateScope(query.ForUpdate)).
			Scopes(productTypeInScope(query.ProductTypeIn)).
			Scopes(productTypeEqualScope(query.ProductType))
	}
}
