package futurescontractdao

import (
	"github.com/jinzhu/gorm"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
)

const (
	table = "futures_contract"
)

//QueryModel lists all queryable columns
type QueryModel struct {
	FutureVarietyCode string
	ContractCode      string
	ContractCodeIn    []string
	ExchangeCode      string
	ExchangeCodeIn    []string
	ContractMonth     string
	Status            string
}

func New(tx *gorm.DB, futuresProduct *models.FuturesContractModel) {
	err := tx.Table(table).
		Create(&futuresProduct).Error

	if err != nil {
		panic(err)
	}
}

func Modify(tx *gorm.DB, product *models.FuturesContractModel) {
	attrs := map[string]interface{}{
		//"last_trade_date": product.LastTradeDate,
		"contract_status": product.ContractStatus,
	}
	err := tx.Table(table).
		Model(models.FuturesContractModel{}).
		Where("contract_code = ?", product.ContractCode).Debug().
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

func Gets(tx *gorm.DB, query *QueryModel) []*models.FuturesContractModel {
	var result []*models.FuturesContractModel
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return make([]*models.FuturesContractModel, 0)
	}

	if err != nil {
		panic(err)
	}

	return result
}

func Get(tx *gorm.DB, query *QueryModel) *models.FuturesContractModel {
	result := &models.FuturesContractModel{}
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

func contractCodeEqualScope(contractCode string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if contractCode != "" {
			return db.Where(table+".contract_code = ?", contractCode)
		}
		return db
	}
}

func contractMonthEqualScope(contractMonth string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if contractMonth != "" {
			return db.Where(table+".month = ?", contractMonth)
		}
		return db
	}

}

func statusEqualScope(status string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != "" {
			return db.Where(table+".contract_status = ?", status)
		}
		return db
	}
}

func futureVarietyCodeEqualScope(futureVarietyCode string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if futureVarietyCode != "" {
			return db.Where(table+".future_variety_code = ?", futureVarietyCode)
		}
		return db
	}
}

func exchangeCodeEqualScope(exchangeCode string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if exchangeCode != "" {
			return db.Where("futures_contract"+".exchange_code = ?", exchangeCode)
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

func contractCodeInScope(symbolIn []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(symbolIn) > 0 {
			return db.Where(table+".contract_code IN (?)", symbolIn)
		}
		return db
	}
}

func queryChain(query *QueryModel) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Scopes(contractCodeEqualScope(query.ContractCode)).
			Scopes(contractMonthEqualScope(query.ContractMonth)).
			Scopes(statusEqualScope(query.Status)).
			Scopes(exchangeCodeEqualScope(query.ExchangeCode)).
			Scopes(exchangeCodeInScope(query.ExchangeCodeIn)).
			Scopes(contractCodeInScope(query.ContractCodeIn)).
			Scopes(futureVarietyCodeEqualScope(query.FutureVarietyCode))
	}
}
