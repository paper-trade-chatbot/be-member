package exchangedao

import (
	"fmt"

	"github.com/jinzhu/gorm"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
	"github.com/spf13/cast"
)

const (
	table = "exchange"
)

//QueryModel lists all queryable columns
type QueryModel struct {
	ExchangeCode   string
	ExchangeCodeIn []string
	ID             string
	Status         string
}

func New(tx *gorm.DB, exchange *models.ExchangeModel) {
	err := tx.Table(table).
		Create(&exchange).Error

	if err != nil {
		panic(err)
	}
}

func Gets(tx *gorm.DB, query *QueryModel) []*models.ExchangeModel {
	var result []*models.ExchangeModel
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return make([]*models.ExchangeModel, 0)
	}

	if err != nil {
		panic(err)
	}

	return result
}

func Get(tx *gorm.DB, query *QueryModel) *models.ExchangeModel {
	result := &models.ExchangeModel{}
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

func Modify(tx *gorm.DB, exchange *models.ExchangeModel) {
	attrs := map[string]interface{}{
		"exchange_status": exchange.ExchangeStatus,
		"trade_time":      exchange.TradeTime,
		"exception_time":  exchange.ExceptionTime,
	}
	err := tx.Table(table).
		Model(models.ExchangeModel{}).
		Where("id = ?", exchange.ID).
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

func exchangeCodeInScope(symbolIn []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(symbolIn) > 0 {
			return db.Where(table+".exchange_code IN (?)", symbolIn)
		}
		return db
	}
}

func idEqualScope(id string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if id != "" {
			return db.Where(table+".id = ?", getId(id))
		}
		return db
	}
}

func statusEqualScope(status string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != "" {
			return db.Where(table+".exchange_status = ?", status)
		}
		return db
	}
}

func queryChain(query *QueryModel) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Scopes(exchangeCodeEqualScope(query.ExchangeCode)).
			Scopes(idEqualScope(query.ID)).
			Scopes(exchangeCodeInScope(query.ExchangeCodeIn)).
			Scopes(statusEqualScope(query.Status))
	}
}

func setId(id int) string {
	return fmt.Sprintf("E%04d", id)
}

func getId(id string) int {
	index := 0
	for i := 1; i < len(id); i++ {
		if string(id[i]) != "0" {
			index = i
			break
		}
	}
	return cast.ToInt(id[index:])
}
