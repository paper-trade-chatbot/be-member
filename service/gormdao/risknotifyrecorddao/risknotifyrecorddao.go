package risknotifyrecorddao

import (
	"time"

	"github.com/jinzhu/gorm"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
	"github.com/paper-trade-chatbot/be-member/util/specification"
)

const (
	table = "risk_notify_record"
)

//QueryModel lists all queryable columns
type QueryModel struct {
	CustNumber   string
	CustNumberIn []string
	OrderBy      []string
	DateStart    string
	DateEnd      string
	Limit        int
	CreatedAfter time.Time
}

func New(tx *gorm.DB, model *models.RiskNotifyRecordModel) {
	err := tx.Table(table).
		Create(&model).Error

	if err != nil {
		panic(err)
	}
}

func Count(tx *gorm.DB, query *QueryModel) int {
	var count int
	err := tx.Table(table).Scopes(queryChain(query)).Count(&count).Error
	if err != nil {
		panic(err)
	}
	return count
}

func Modify(tx *gorm.DB, riskRecord *models.RiskNotifyRecordModel) {
	attrs := map[string]interface{}{}
	err := tx.Table(table).
		Model(models.FuturesOrderModel{}).
		Where("id = ?", riskRecord.ID).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

func GetsWithPagination(tx *gorm.DB, query *QueryModel, paginate *specification.PaginationStruct) ([]*models.RiskNotifyRecordModel, int32) {

	paginate.TableName = table
	var rows []*models.RiskNotifyRecordModel
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

func Gets(tx *gorm.DB, query *QueryModel) []*models.RiskNotifyRecordModel {
	var result []*models.RiskNotifyRecordModel
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

func Get(tx *gorm.DB, query *QueryModel) *models.RiskNotifyRecordModel {
	result := &models.RiskNotifyRecordModel{}
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

func Delete(tx *gorm.DB, clOrderId string) {
	err := tx.Table(table).Where("id = ?", clOrderId).Delete(&models.FuturesPositionModel{}).Error

	if err != nil {
		panic(err)
	}
}

func custNumberEqualScope(custNumber string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if custNumber != "" {
			return db.Where(table+".cust_number = ?", custNumber)
		}
		return db
	}
}

func custNumberInScope(custNumberIn []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(custNumberIn) > 0 {
			return db.Where(table+".cust_number IN (?)", custNumberIn)
		}
		return db
	}
}

func createdAfterScope(createAfter time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if createAfter.IsZero() {
			return db
		}
		return db.Where(table+".created_at > ?", createAfter)
	}
}

func createdAtBetweenScope(startTime, endTime string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if startTime != "" && endTime != "" {
			return db.Where(table+".created_at BETWEEN ? AND ?", startTime, endTime)
		}
		return db
	}
}

func orderByScope(orderBy []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(orderBy) != 0 {
			order := orderBy[0]
			for _, o := range orderBy[1:] {
				order = order + ", " + o
			}
			return db.Order(order)
		}
		return db
	}
}

func limitScope(limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if limit <= 0 {
			limit = 100
		}
		return db.Limit(limit)
	}
}

func queryChain(query *QueryModel) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Scopes(custNumberEqualScope(query.CustNumber)).
			Scopes(custNumberInScope(query.CustNumberIn)).
			Scopes(createdAfterScope(query.CreatedAfter)).
			Scopes(createdAtBetweenScope(query.DateStart, query.DateEnd)).
			Scopes(orderByScope(query.OrderBy)).
			Scopes(limitScope(query.Limit))

	}
}
