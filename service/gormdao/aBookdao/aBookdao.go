package aBookdao

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/now"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
	"github.com/paper-trade-chatbot/be-member/util/specification"
)

const (
	table = "a_book"
)

//QueryModel lists all queryable columns
type QueryModel struct {
	ProductCode    string
	ExchangeCode   string
	ProductType    int32
	OrderBy        []string
	OrderDirection []string
	DateStart      time.Time
	DateEnd        time.Time
}

func New(tx *gorm.DB, aBook *models.ABookModel) {

	err := tx.Table(table).
		Create(&aBook).Error

	if err != nil {
		panic(err)
	}
}

func GetFirstRecord(tx *gorm.DB) *models.ABookModel {
	result := &models.ABookModel{}
	err := tx.Table(table).
		Order("created_at asc").Limit(1).
		Scan(result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}

	if err != nil {
		panic(err)
	}
	return result
}

func Gets(tx *gorm.DB, query *QueryModel) []*models.ABookModel {
	var result []*models.ABookModel
	err := tx.Table(table).
		Select(
			"id AS id," +
				"exchange_code AS exchange_code," +
				"product_code AS product_code, " +
				"product_type AS product_type, " +

				"buy_position_qty AS buy_position_qty," +
				"sell_position_qty AS sell_position_qty," +

				"current_price AS current_price," +
				"contract AS contract," +
				"exchange_rate AS exchange_rate," +
				"eod_price AS eod_price," +

				"hedge_pnl AS hedge_pnl, " +
				"previous_floating_pnl AS previous_floating_pnl, " +
				"today_floating_pnl AS today_floating_pnl, " +
				"previous_realized_pnl AS previous_realized_pnl, " +
				"today_realized_pnl AS today_realized_pnl," +

				"created_at AS created_at," +
				"updated_at AS updated_at," +

				"(buy_position_qty-sell_position_qty) AS net_position," +
				"(buy_position_qty+sell_position_qty) AS gross_position," +
				"((buy_position_qty-sell_position_qty)*current_price*contract*exchange_rate) AS net_exposure," +
				"(hedge_pnl+previous_floating_pnl+today_floating_pnl+previous_realized_pnl+today_realized_pnl) AS total_pnl").
		Scopes(queryChain(query)).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return make([]*models.ABookModel, 0)
	}

	if err != nil {
		panic(err)
	}

	return result
}

func GetsWithPagination(tx *gorm.DB, query *QueryModel, paginate *specification.PaginationStruct) ([]*models.ABookModel, int32) {
	paginate.TableName = table
	var rows []*models.ABookModel
	var count int32 = 0
	err := tx.Table(table).
		Select(
			"id AS id," +
				"exchange_code AS exchange_code," +
				"product_code AS product_code, " +
				"product_type AS product_type, " +

				"buy_position_qty AS buy_position_qty," +
				"sell_position_qty AS sell_position_qty," +

				"current_price AS current_price," +
				"contract AS contract," +
				"exchange_rate AS exchange_rate," +
				"eod_price AS eod_price," +

				"hedge_pnl AS hedge_pnl, " +
				"previous_floating_pnl AS previous_floating_pnl, " +
				"today_floating_pnl AS today_floating_pnl, " +
				"previous_realized_pnl AS previous_realized_pnl, " +
				"today_realized_pnl AS today_realized_pnl," +

				"created_at AS created_at," +
				"updated_at AS updated_at," +

				"(buy_position_qty-sell_position_qty) AS net_position," +
				"(buy_position_qty+sell_position_qty) AS gross_position," +
				"((buy_position_qty-sell_position_qty)*current_price*contract*exchange_rate) AS net_exposure," +
				"(hedge_pnl+previous_floating_pnl+today_floating_pnl+previous_realized_pnl+today_realized_pnl) AS total_pnl").
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

func SumForPagination(tx *gorm.DB, query *QueryModel) *models.ABookModel {
	result := &models.ABookModel{}
	err := tx.Table(table).
		Select(
			"sum(buy_position_qty) AS buy_position_qty," +
				"sum(sell_position_qty) AS sell_position_qty," +
				"sum(hedge_pnl) AS hedge_pnl, " +
				"sum(previous_floating_pnl) AS previous_floating_pnl, " +
				"sum(today_floating_pnl) AS today_floating_pnl, " +
				"sum(previous_realized_pnl) AS previous_realized_pnl, " +
				"sum(today_realized_pnl) AS today_realized_pnl," +

				"sum(buy_position_qty-sell_position_qty) AS net_position," +
				"sum(buy_position_qty+sell_position_qty) AS gross_position," +
				"sum(abs(buy_position_qty-sell_position_qty)*current_price*contract*exchange_rate) AS net_exposure," +
				"sum(hedge_pnl+previous_floating_pnl+today_floating_pnl+previous_realized_pnl+today_realized_pnl) AS total_pnl").
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

func Count(tx *gorm.DB, query *QueryModel) int32 {
	var count int32
	err := tx.Table(table).Scopes(queryChain(query)).Count(&count).Error
	if err != nil {
		panic(err)
	}
	return count
}

func Get(tx *gorm.DB, query *QueryModel) *models.ABookModel {
	result := &models.ABookModel{}
	err := tx.Table(table).
		Select(
			"id AS id," +
				"exchange_code AS exchange_code," +
				"product_code AS product_code, " +
				"product_type AS product_type, " +

				"buy_position_qty AS buy_position_qty," +
				"sell_position_qty AS sell_position_qty," +

				"current_price AS current_price," +
				"contract AS contract," +
				"exchange_rate AS exchange_rate," +
				"eod_price AS eod_price," +

				"hedge_pnl AS hedge_pnl, " +
				"previous_floating_pnl AS previous_floating_pnl, " +
				"today_floating_pnl AS today_floating_pnl, " +
				"previous_realized_pnl AS previous_realized_pnl, " +
				"today_realized_pnl AS today_realized_pnl," +

				"created_at AS created_at," +
				"updated_at AS updated_at," +

				"(buy_position_qty-sell_position_qty) AS net_position," +
				"(buy_position_qty+sell_position_qty) AS gross_position," +
				"((buy_position_qty-sell_position_qty)*current_price*contract*exchange_rate) AS net_exposure," +
				"(hedge_pnl+previous_floating_pnl+today_floating_pnl+previous_realized_pnl+today_realized_pnl) AS total_pnl").
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

func SumGroupByProductCode(tx *gorm.DB, query *QueryModel) []*models.ABookModel {
	var result []*models.ABookModel
	err := tx.Table(table).
		Select(
			"exchange_code, product_code, product_type, " +
				"sum(buy_position_qty) AS buy_position_qty," +
				"sum(sell_position_qty) AS sell_position_qty," +
				"sum(hedge_pnl) AS hedge_pnl, " +
				"sum(previous_floating_pnl) AS previous_floating_pnl, " +
				"sum(today_floating_pnl) AS today_floating_pnl, " +
				"sum(previous_realized_pnl) AS previous_realized_pnl, " +
				"sum(today_realized_pnl) AS today_realized_pnl," +
				"sum(hedge_pnl+previous_floating_pnl+today_floating_pnl+previous_realized_pnl+today_realized_pnl) AS total_pnl").
		Scopes(queryChain(query)).
		Group("exchange_code, product_code, product_type").
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}

	if err != nil {
		panic(err)
	}

	return result
}

func SumPNL(tx *gorm.DB, query *QueryModel) []*models.ABookModel {
	var result []*models.ABookModel
	err := tx.Table(table).
		Select(
			"sum(hedge_pnl) AS hedge_pnl, " +
				"sum(previous_floating_pnl) AS previous_floating_pnl, " +
				"sum(today_floating_pnl) AS today_floating_pnl, " +
				"sum(previous_realized_pnl) AS previous_realized_pnl, " +
				"sum(today_realized_pnl) AS today_realized_pnl, " +
				"sum(hedge_pnl+previous_floating_pnl+today_floating_pnl+previous_realized_pnl+today_realized_pnl) AS total_pnl").
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

func Modify(tx *gorm.DB, aBook *models.ABookModel, attrs map[string]interface{}) {

	err := tx.Table(table).
		Model(models.ABookModel{}).
		Where("id = ?", aBook.ID).
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

func productTypeEqualScope(productType int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if productType != 0 {
			return db.Where(table+".product_type = ?", productType)
		}
		return db
	}
}

func orderByScope(orderBy []string, orderDirection []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(orderBy) != len(orderDirection) {
			return db
		}
		if len(orderBy) != 0 {
			order := ""
			for idx, o := range orderBy {
				if idx == 0 {
					order = replaceOrderBy(o, orderDirection[idx])
				} else {
					order = order + ", " + replaceOrderBy(o, orderDirection[idx])
				}
			}
			return db.Order(order)
		}
		return db
	}
}

func replaceOrderBy(in, direction string) string {
	orderItem := ""
	switch in {
	case "net_position":
		orderItem = "abs(buy_position_qty-sell_position_qty)"
	case "net_exposure":
		orderItem = "abs((buy_position_qty-sell_position_qty)*current_price*contract*exchange_rate)"
	case "hedge_pnl":
		orderItem = "abs(hedge_pnl)"
	case "floating_pnl":
		orderItem = "abs(previous_floating_pnl+today_floating_pnl)"
	case "realized_pnl":
		orderItem = "abs(previous_realized_pnl+today_realized_pnl)"
	case "total_pnl":
		orderItem = "abs(previous_floating_pnl+today_floating_pnl)+abs(previous_realized_pnl+today_realized_pnl)+abs(hedge_pnl)"
	default:
		return ""
	}

	return orderItem + " " + direction
}

func dateBetweenEqualScope(dateStart, dateEnd time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !dateStart.IsZero() && !dateEnd.IsZero() {
			dateS := now.New(dateStart).BeginningOfDay()
			dateE := now.New(dateEnd).EndOfDay()
			return db.Where(table+".created_at BETWEEN ? AND ?", dateS, dateE)
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
			Scopes(orderByScope(query.OrderBy, query.OrderDirection)).
			Scopes(dateBetweenEqualScope(query.DateStart, query.DateEnd))
	}
}
