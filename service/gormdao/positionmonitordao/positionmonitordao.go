package positionmonitordao

import (
	"github.com/jinzhu/gorm"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
	"github.com/shopspring/decimal"
)

const (
	table = "position_monitor"
)

//QueryModel lists all queryable columns
type QueryModel struct {
	ProductCode    string
	ExchangeCode   string
	ProductType    int32
	ProductCodeArr []string
}

func New(tx *gorm.DB, positionMonitor *models.PositionMonitorModel) {

	if positionMonitor.ClientBuyQty.Add(positionMonitor.CompanyBuyQty).GreaterThan(decimal.Zero) {
		positionMonitor.TotalBuyAvgPrice = positionMonitor.ClientBuyAvgPrice.Mul(positionMonitor.ClientBuyQty).Add(positionMonitor.CompanyBuyAvgPrice.Mul(positionMonitor.CompanyBuyQty)).Div(positionMonitor.ClientBuyQty.Add(positionMonitor.CompanyBuyQty))
	} else {
		positionMonitor.TotalBuyAvgPrice = decimal.Zero
	}

	if positionMonitor.ClientSellQty.Add(positionMonitor.CompanySellQty).GreaterThan(decimal.Zero) {
		positionMonitor.TotalSellAvgPrice = positionMonitor.ClientSellAvgPrice.Mul(positionMonitor.ClientSellQty).Add(positionMonitor.CompanySellAvgPrice.Mul(positionMonitor.CompanySellQty)).Div(positionMonitor.ClientSellQty.Add(positionMonitor.CompanySellQty))
	} else {
		positionMonitor.TotalSellAvgPrice = decimal.Zero
	}

	err := tx.Table(table).
		Create(&positionMonitor).Error

	if err != nil {
		panic(err)
	}
}

func Gets(tx *gorm.DB, query *QueryModel) []*models.PositionMonitorModel {
	var result []*models.PositionMonitorModel
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return make([]*models.PositionMonitorModel, 0)
	}

	if err != nil {
		panic(err)
	}

	return result
}

func Get(tx *gorm.DB, query *QueryModel) *models.PositionMonitorModel {
	result := &models.PositionMonitorModel{}
	err := tx.Table(table).
		Scopes(queryChain(query)).
		First(result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}

	if err != nil {
		panic(err)
	}
	return result
}

func Modify(tx *gorm.DB, exchange *models.PositionMonitorModel) {
	attrs := map[string]interface{}{
		"client_buy_avg_price":  exchange.ClientBuyAvgPrice,
		"client_buy_qty":        exchange.ClientBuyQty,
		"client_sell_avg_price": exchange.ClientSellAvgPrice,
		"client_sell_qty":       exchange.ClientSellQty,

		"company_buy_avg_price":  exchange.CompanyBuyAvgPrice,
		"company_buy_qty":        exchange.CompanyBuyQty,
		"company_sell_avg_price": exchange.CompanySellAvgPrice,
		"company_sell_qty":       exchange.CompanySellQty,
	}
	err := tx.Table(table).
		Model(models.PositionMonitorModel{}).
		Where("exchange_code = ? AND product_code = ?", exchange.ExchangeCode, exchange.ProductCode).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

func exchangeCodeEqualScope(symbol string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if symbol != "" {
			return db.Where(table+".exchange_code = ?", symbol)
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

func productCodeInScope(productCodeArr []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(productCodeArr) > 0 {
			return db.Where(table+".product_code IN (?)", productCodeArr)
		}
		return db
	}
}

func productTypeEqualScope(productType int32) func(db *gorm.DB) *gorm.DB {
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
			Scopes(exchangeCodeEqualScope(query.ExchangeCode)).
			Scopes(productCodeEqualScope(query.ProductCode)).
			Scopes(productTypeEqualScope(query.ProductType)).
			Scopes(productCodeInScope(query.ProductCodeArr))
	}
}
