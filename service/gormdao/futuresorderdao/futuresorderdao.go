package futuresorderdao

import (
	"time"

	"github.com/jinzhu/gorm"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
	"github.com/paper-trade-chatbot/be-member/util/specification"
)

const (
	table = "futures_order"
)

//QueryModel lists all queryable columns
type QueryModel struct {
	ClOrderId             string
	SysOrderId            string
	CustNumber            string
	CustNumberIn          []string
	TradingMemberNumber   string
	TradingMemberNumberIn []string
	ContractCodeIn        []string
	Side                  string
	OpenClose             string
	OrderType             int
	OrderTypeIn           []int
	OrderStatus           int
	OrderBy               []string
	DateStart             string
	DateEnd               string
	CreatedAfter          time.Time
}

func New(tx *gorm.DB, model *models.FuturesOrderModel) {
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

func Modify(tx *gorm.DB, order *models.FuturesOrderModel) {
	attrs := map[string]interface{}{
		"sys_order_id":   order.SysOrderID,
		"open_close":     order.OpenClose,
		"order_type":     order.OrderType,
		"order_price":    order.OrderPrice,
		"order_qty":      order.OrderQty,
		"commit_price":   order.CommitPrice,
		"commit_qty":     order.CommitQty,
		"order_status":   order.OrderStatus,
		"settled_profit": order.SettledProfit,
	}
	err := tx.Table(table).
		Model(models.FuturesOrderModel{}).
		Where("cl_order_id = ?", order.ClOrderID).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

func GetsWithPagination(tx *gorm.DB, query *QueryModel, paginate *specification.PaginationStruct) ([]*models.FuturesOrderModel, int32) {

	paginate.TableName = table
	var rows []*models.FuturesOrderModel
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

func Gets(tx *gorm.DB, query *QueryModel) []*models.FuturesOrderModel {
	var result []*models.FuturesOrderModel
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

func Get(tx *gorm.DB, query *QueryModel) *models.FuturesOrderModel {
	result := &models.FuturesOrderModel{}
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
	err := tx.Table(table).Where("cl_order_id=?", clOrderId).Delete(&models.FuturesPositionModel{}).Error

	if err != nil {
		panic(err)
	}
}

func clOrderIdEqualScope(clOrderId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if clOrderId != "" {
			return db.Where(table+".cl_order_id = ?", clOrderId)
		}
		return db
	}
}

func orderStatusEqualScope(status int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != 0 {
			return db.Where(table+".order_status = ?", status)
		}
		return db
	}
}

func tradingMemberNumberEqualScope(account string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if account != "" {
			return db.Where(table+".trading_member_number = ?", account)
		}
		return db
	}
}

func tradingMemberNumberInScope(memberIn []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(memberIn) > 0 {
			return db.Where(table+".trading_member_number IN (?)", memberIn)
		}
		return db
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

func contractCodeInScope(contractCodeIn []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(contractCodeIn) > 0 {
			return db.Where(table+".cust_number IN (?)", contractCodeIn)
		}
		return db
	}
}

func sideEqualScope(buySell string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if buySell != "" {
			return db.Where(table+".Side = ?", buySell)
		}
		return db
	}
}

func systemOrderIdEqualScope(systemOrderId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if systemOrderId != "" {
			return db.Where(table+".sys_order_id = ?", systemOrderId)
		}
		return db
	}
}

func initialMarginTotalNotZero(flag bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if flag {
			return db.Where(table+".initial_margin_total != ?", 0)
		}
		return db
	}
}

func productNameEqualScope(productName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if productName != "" {
			return db.Where("futures_product"+".name = ?", productName)
		}
		return db
	}
}

func createTimeBetweenScope(startTime, endTime string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if startTime == "" || endTime == "" {
			return db
		}
		return db.Where(table+".created_at BETWEEN ? AND ?", startTime, endTime)
	}
}

func updateTimeBetweenScope(startTime, endTime string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if startTime == "" || endTime == "" {
			return db
		}
		return db.Where(table+".updated_at BETWEEN ? AND ?", startTime, endTime)
	}
}

func OrderTypeInScope(orderTypeIn []int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(orderTypeIn) > 0 {
			return db.Where(table+".order_type IN (?)", orderTypeIn)
		}
		return db
	}
}

func orderTypeScope(orderType int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if orderType != 0 {
			return db.Where(table+".order_type = ?", orderType)
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

func queryChain(query *QueryModel) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Scopes(orderStatusEqualScope(query.OrderStatus)).
			Scopes(tradingMemberNumberEqualScope(query.TradingMemberNumber)).
			Scopes(tradingMemberNumberInScope(query.TradingMemberNumberIn)).
			Scopes(custNumberEqualScope(query.CustNumber)).
			Scopes(custNumberInScope(query.CustNumberIn)).
			Scopes(contractCodeInScope(query.ContractCodeIn)).
			Scopes(systemOrderIdEqualScope(query.SysOrderId)).
			Scopes(sideEqualScope(query.Side)).
			Scopes(OrderTypeInScope(query.OrderTypeIn)).
			Scopes(orderTypeScope(query.OrderType)).
			Scopes(clOrderIdEqualScope(query.ClOrderId)).
			Scopes(createdAfterScope(query.CreatedAfter)).
			Scopes(createdAtBetweenScope(query.DateStart, query.DateEnd)).
			Scopes(orderByScope(query.OrderBy))
	}
}
