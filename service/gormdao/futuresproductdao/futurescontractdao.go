package futuresproductdao

import (
	"github.com/jinzhu/gorm"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
)

const (
	table = "futures_product"
)

//QueryModel lists all queryable columns
type QueryModel struct {
	FuturesVarietyCode     string
	FuturesVarietyCodeID   int32
	FuturesVarietyCodeIDIn []int32
	ProductCode            string
	ProductCodeIn          []string
	ExchangeCode           string
	ExchangeCodeIn         []string
	ContractMonth          string
	ProductPinYinIN        string
	Status                 string
	Display                string
	Pattern                string
	Orderby                []string
}

func New(tx *gorm.DB, futuresProduct *models.FuturesProductModel) error {
	err := tx.Table(table).
		Create(&futuresProduct).Error

	if err != nil {
		return err
	}
	return nil
}

func Modify(tx *gorm.DB, product *models.FuturesProductModel) {
	attrs := map[string]interface{}{
		"margin":         product.Margin,
		"product_status": product.ProductStatus,
		"display":        product.Display,
	}
	err := tx.Table(table).
		Model(models.FuturesProductModel{}).
		Where("exchange_code = ? and product_code = ?", product.ExchangeCode, product.ProductCode).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}

	// 同步 futures_contract, futures_product 的改動
	attrs["contract_status"] = attrs["product_status"]
	delete(attrs, "product_status")
	err = tx.Table("futures_contract").
		Model(models.FuturesContractModel{}).
		Where("product_code = ?", product.ProductCode).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

func Gets(tx *gorm.DB, query *QueryModel) []*models.FuturesProductModel {
	var result []*models.FuturesProductModel
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return make([]*models.FuturesProductModel, 0)
	}

	if err != nil {
		panic(err)
	}

	return result
}

func Get(tx *gorm.DB, query *QueryModel) *models.FuturesProductModel {
	result := &models.FuturesProductModel{}
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

func productPinYinEqualScope(ProductPinYinIN string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if ProductPinYinIN != "" {
			return db.Where(table+".product_pinyin like  ? or product_code like ? or product_name like ?", "%"+ProductPinYinIN+"%", "%"+ProductPinYinIN+"%", "%"+ProductPinYinIN+"%")
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

func futuresVarietyCodeEqualScope(futureVarietyCode string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if futureVarietyCode != "" {
			return db.Where(table+".futures_variety_code = ?", futureVarietyCode)
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

func exchangeCodeInScope(symbolIn []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(symbolIn) > 0 && symbolIn[0] != "" {
			return db.Where(table+".exchange_code IN (?)", symbolIn)
		}
		return db
	}
}

func futuresVarietyCodeInScope(codes []int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(codes) > 0 && codes[0] != 0 {
			return db.Where(table+".futures_variety_code_id IN (?)", codes)
		}
		return db
	}
}

func productCodeInScope(symbolIn []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(symbolIn) > 0 && symbolIn[0] != "" {
			return db.Where(table+".product_code IN (?)", symbolIn)
		}
		return db
	}
}

func productCodeOrNameOrPinyiLikeScope(producrCode string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if producrCode != "" {
			return db.Where(table+".product_code like ? or product_pinyin LIKE ?", "%"+producrCode+"%", "%"+producrCode+"%")
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

func queryChain(query *QueryModel) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Scopes(productCodeEqualScope(query.ProductCode)).
			Scopes(statusEqualScope(query.Status)).
			Scopes(exchangeCodeEqualScope(query.ExchangeCode)).
			Scopes(exchangeCodeInScope(query.ExchangeCodeIn)).
			Scopes(productCodeInScope(query.ProductCodeIn)).
			Scopes(productPinYinEqualScope(query.ProductPinYinIN)).
			Scopes(futuresVarietyCodeEqualScope(query.FuturesVarietyCode)).
			Scopes(productCodeOrNameOrPinyiLikeScope(query.Pattern)).
			Scopes(futuresVarietyCodeInScope(query.FuturesVarietyCodeIDIn)).
			Scopes(orderByScope(query.Orderby)).
			Scopes(displayEqualScope(query.Display))

	}
}
