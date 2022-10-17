package cryptoproductdao

import (
	"github.com/jinzhu/gorm"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
)

const (
	table = "crypto_product"
)

//QueryModel lists all queryable columns
type QueryModel struct {
	ProductCode    string
	ProductCodeIn  []string
	ExchangeCode   string
	ExchangeCodeIn []string
	Status         string
	Display        string
	Pattern        string
}

func New(tx *gorm.DB, futuresProduct *models.CryptoProductModel) error {
	err := tx.Table(table).
		Create(&futuresProduct).Error

	if err != nil {
		return err
	}
	return nil
}

func Modify(tx *gorm.DB, product *models.CryptoProductModel) {
	attrs := map[string]interface{}{
		"product_status": product.ProductStatus,
		"display":        product.Display,
	}
	err := tx.Table(table).
		Model(models.CryptoProductModel{}).
		Where("exchange_code = ? and product_code = ?", product.ExchangeCode, product.ProductCode).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

func Gets(tx *gorm.DB, query *QueryModel) []*models.CryptoProductModel {
	var result []*models.CryptoProductModel
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return make([]*models.CryptoProductModel, 0)
	}

	if err != nil {
		panic(err)
	}

	return result
}

func Get(tx *gorm.DB, query *QueryModel) *models.CryptoProductModel {
	result := &models.CryptoProductModel{}
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

func productCodeEqualScope(producrCode string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if producrCode != "" {
			return db.Where(table+".product_code = ?", producrCode)
		}
		return db
	}
}

func productCodeOrNameOrPinyiLikeScope(producrCode string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if producrCode != "" {
			return db.Where(table+".product_code like ? or product_name_en LIKE ? or product_pinyin LIKE ?", "%"+producrCode+"%", "%"+producrCode+"%", "%"+producrCode+"%")
		}
		return db
	}
}

func statusEqualScope(status string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != "" {
			return db.Where(table+".product_status = ?", status)
		}
		return db
	}
}

func displayEqualScope(display string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if display != "" {
			return db.Where(table+".display = ?", display)
		}
		return db
	}
}

// func futuresVarietyCodeEqualScope(futureVarietyCode string) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		if futureVarietyCode != "" {
// 			return db.Where(table+".futures_variety_code = ?", futureVarietyCode)
// 		}
// 		return db
// 	}
// }

func exchangeCodeEqualScope(exchangeCode string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if exchangeCode != "" {
			return db.Where(table+".exchange_code = ?", exchangeCode)
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

func productCodeInScope(symbolIn []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(symbolIn) > 0 {
			return db.Where(table+".product_code IN (?)", symbolIn)
		}
		return db
	}
}

func queryChain(query *QueryModel) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Scopes(productCodeEqualScope(query.ProductCode)).
			Scopes(statusEqualScope(query.Status)).
			Scopes(exchangeCodeEqualScope(query.ExchangeCode)).
			Scopes(exchangeCodeInScope(query.ExchangeCodeIn)).
			Scopes(productCodeInScope(query.ProductCodeIn)).
			Scopes(productCodeOrNameOrPinyiLikeScope(query.Pattern)).
			Scopes(displayEqualScope(query.Display))
	}
}
