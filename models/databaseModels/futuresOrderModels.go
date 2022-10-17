package databaseModels

import "time"

type FuturesOrderModel struct {
	ClOrderID                   string     `gorm:"column:cl_order_id"`
	SysOrderID                  string     `gorm:"column:sys_order_id"`
	TradingMemberNumber         string     `gorm:"column:trading_member_number"`
	CustNumber                  string     `gorm:"column:cust_number"`
	SeatID                      string     `gorm:"column:seat_id"`
	ExchangeCode                string     `gorm:"column:exchange_code"`
	ContractCode                string     `gorm:"column:contract_code"`
	Side                        int        `gorm:"column:side"`
	OrderType                   int        `gorm:"column:order_type"`
	OpenClose                   int        `gorm:"column:open_close"`
	OrderPrice                  float64    `gorm:"column:order_price"`
	OrderQty                    float64    `gorm:"column:order_qty"`
	CommitPrice                 float64    `gorm:"column:commit_price"`
	CommitQty                   float64    `gorm:"column:commit_qty"`
	TimeInForce                 int        `gorm:"column:time_in_force"`
	OrderStatus                 int        `gorm:"column:order_status"`
	SettledProfit               float64    `gorm:"column:settled_profit"`
	ForcedLiquidationReasonCode int        `gorm:"column:forced_liquidation_reason_code"`
	CommitDate                  *time.Time `gorm:"column:commit_date"`
	TradingDate                 *time.Time `gorm:"column:trading_date"`
	ForceDropFlag               int        `gorm:"column:force_drop_flag"`
	IP                          string     `gorm:"column:ip"`
	MAC                         string     `gorm:"column:mac"`
	OrderSystem                 string     `gorm:"column:order_system"`
	OperatorNo                  string     `gorm:"column:operator_no"`
	CreatedAt                   time.Time  `gorm:"column:created_at"`
	UpdatedAt                   time.Time  `gorm:"column:updated_at"`
}

func (order *FuturesOrderModel) Insert(queue *[]*FuturesOrderModel) bool {
	switch order.OrderType {
	case 1: //market
		*queue = append(*queue, order)
	case 3: //stop
		*queue = append(*queue, order)
	case 2: // limit
		var insertIndex int
		insertIndex = -1
		if order.Side == 1 { //buy
			//   299  298 大於改小於
			for i, queuedOrder := range *queue {
				if order.OrderPrice > queuedOrder.OrderPrice {
					insertIndex = i
					break
				} else if order.OrderPrice == queuedOrder.OrderPrice {
					if order.CreatedAt.Before(queuedOrder.CreatedAt) {
						insertIndex = i
						break
					}
				}
			}
		} else {
			for i, queuedOrder := range *queue {
				if order.OrderPrice < queuedOrder.OrderPrice {
					insertIndex = i
					break
				} else if order.OrderPrice == queuedOrder.OrderPrice {
					if order.CreatedAt.Before(queuedOrder.CreatedAt) {
						insertIndex = i
						break
					}
				}
			}
		}

		if insertIndex == -1 {
			*queue = append(*queue, order)
		} else if insertIndex == 0 {
			var queue2 []*FuturesOrderModel
			queue2 = append(queue2, order)
			*queue = append(queue2, (*queue)[insertIndex:]...)
		} else {
			*queue = append(*queue, order)
			copy((*queue)[insertIndex+1:], (*queue)[insertIndex:])
			(*queue)[insertIndex] = order
		}
	}
	return true
}

func (order *FuturesOrderModel) Remove(queue *[]*FuturesOrderModel) bool {
	var removeIndex int
	var removeStatus bool

	for i, queuedOrder := range *queue {
		// check order identity
		if order.ClOrderID == queuedOrder.ClOrderID {
			removeIndex = i
			removeStatus = true
			break
		}
	}
	if removeStatus {
		*queue = append((*queue)[:removeIndex], (*queue)[removeIndex+1:]...)
		return true
	} else {
		return false
	}
}
