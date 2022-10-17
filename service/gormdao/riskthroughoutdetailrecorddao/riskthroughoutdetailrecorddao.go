package riskthroughoutdetailrecorddao

import (
	"github.com/jinzhu/gorm"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
	"github.com/paper-trade-chatbot/be-member/util/specification"
)

const (
	table = "risk_throughout_detail_record"
)

//QueryModel lists all queryable columns
type QueryModel struct {
	ProductCode  string
	ExchangeCode string
	OrderBy      []string
	DateStart    string
	DateEnd      string
	Limit        int
	ProductType  int32
	RecordId     string
}

func New(tx *gorm.DB, positionMonitor *models.RiskThroughoutDetailRecordModel) {

	positionMonitor.ThroughoutLoss = (positionMonitor.AvgPrice.Sub(positionMonitor.ThroughoutAvgPrice)).Mul(positionMonitor.ThroughoutPositionQty)
	err := tx.Table(table).
		Create(&positionMonitor).Error

	if err != nil {
		panic(err)
	}
}

func Gets(tx *gorm.DB, query *QueryModel) []*models.RiskThroughoutDetailRecordModel {
	var result []*models.RiskThroughoutDetailRecordModel
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return make([]*models.RiskThroughoutDetailRecordModel, 0)
	}

	if err != nil {
		panic(err)
	}

	return result
}

func Get(tx *gorm.DB, query *QueryModel) *models.RiskThroughoutDetailRecordModel {
	result := &models.RiskThroughoutDetailRecordModel{}
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

func GetsWithPagination(tx *gorm.DB, query *QueryModel, paginate *specification.PaginationStruct) ([]*models.RiskThroughoutDetailRecordModel, int32) {

	paginate.TableName = table
	var rows []*models.RiskThroughoutDetailRecordModel
	var count int32 = 0
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Count(&count).
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

func Modify(tx *gorm.DB, exchange *models.RiskThroughoutDetailRecordModel) {
	attrs := map[string]interface{}{}
	err := tx.Table(table).
		Model(models.RiskThroughoutDetailRecordModel{}).
		Where("id = ? ", exchange.ID).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

// Delete a row
func Delete(tx *gorm.DB, record models.RiskThroughoutDetailRecordModel) {
	err := tx.Table(table).
		Where("id = ?", record.ID).
		Delete(models.RiskThroughoutDetailRecordModel{}).Error

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

func productTypeEqualScope(productType int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if productType != 0 {
			return db.Where(table+".product_type = ?", productType)
		}
		return db
	}
}

func recordIdEqualScope(recordId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if recordId != "" {
			return db.Where(table+".record_id = ?", recordId)
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
			Scopes(recordIdEqualScope(query.RecordId))
	}
}
