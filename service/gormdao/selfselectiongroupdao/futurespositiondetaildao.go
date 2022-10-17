package selfselectiongroupdao

import (
	"github.com/jinzhu/gorm"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
)

const (
	table = "self_selection_group"
)

//QueryModel lists all queryable columns
type QueryModel struct {
	CustNumber   string
	CustNumberIn []string
}

func New(tx *gorm.DB, selfSelectionGroup *models.SelfSelectionGroupModel) {
	err := tx.Table(table).
		Create(&selfSelectionGroup).Error

	if err != nil {
		panic(err)
	}
}

func Delete(tx *gorm.DB, order *models.SelfSelectionGroupModel) {
	err := tx.Table(table).
		Where("cust_number = ?", order.CustNumber).
		Delete(&models.SelfSelectionGroupModel{}).Error

	if err != nil {
		panic(err)
	}
}

func Modify(tx *gorm.DB, order *models.SelfSelectionGroupModel) {
	attrs := map[string]interface{}{
		"group": order.Group,
	}
	err := tx.Table(table).
		Model(models.SelfSelectionGroupModel{}).
		Where("cust_number = ?", order.CustNumber).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

func Gets(tx *gorm.DB, query *QueryModel) []*models.SelfSelectionGroupModel {
	var result []*models.SelfSelectionGroupModel
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

func Get(tx *gorm.DB, query *QueryModel) *models.SelfSelectionGroupModel {
	result := &models.SelfSelectionGroupModel{}
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

func queryChain(query *QueryModel) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Scopes(custNumberEqualScope(query.CustNumber)).
			Scopes(custNumberInScope(query.CustNumberIn))
	}
}
