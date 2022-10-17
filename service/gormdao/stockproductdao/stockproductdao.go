package stockproductdao

import (
	"github.com/jinzhu/gorm"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
)

const (
	table = "stock_product"
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
	Orderby        []string
}

func New(tx *gorm.DB, futuresProduct *models.StockProductModel) error {
	err := tx.Table(table).
		Create(&futuresProduct).Error

	if err != nil {
		return err
	}
	return nil
}

func Modify(tx *gorm.DB, product *models.StockProductModel) {
	attrs := map[string]interface{}{
		"product_status": product.ProductStatus,
		"display":        product.Display,
	}
	err := tx.Table(table).
		Model(models.StockProductModel{}).
		Where("exchange_code = ? and product_code = ?", product.ExchangeCode, product.ProductCode).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

func Gets(tx *gorm.DB, query *QueryModel) []*models.StockProductModel {
	var result []*models.StockProductModel
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return make([]*models.StockProductModel, 0)
	}

	if err != nil {
		panic(err)
	}

	return result
}

func Get(tx *gorm.DB, query *QueryModel) *models.StockProductModel {
	result := &models.StockProductModel{}
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

func productCodeInScope(symbolIn []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(symbolIn) > 0 && symbolIn[0] != "" {
			return db.Where(table+".product_code IN (?)", symbolIn)
		}
		return db
	}
}

func productPinYinEqualScope(ProductPinYinIN string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if ProductPinYinIN != "" {
			return db.Where(table+".product_pinyin like ?", ProductPinYinIN)
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
			Scopes(productCodeOrNameOrPinyiLikeScope(query.Pattern)).
			Scopes(orderByScope(query.Orderby)).
			Scopes(displayEqualScope(query.Display))
	}
}
