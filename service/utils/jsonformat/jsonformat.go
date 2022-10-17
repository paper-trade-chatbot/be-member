package jsonformat

import (
	"time"

	"github.com/shopspring/decimal"
)

type QuoteJSONTemp struct {
	Id         string `json:"id"`
	ProductId  string `json:"productId"`
	ClosePrice string `json:"closePrice"`
	// Bid        string `json:"buyPrice"`
	// Ask        string `json:"sellPrice"`
	// PriorSettle string `json:"priorSettle"`
	// Time        int64  `json:"time"`
}

type QuoteJSON struct {
	ProductId  string          `json:"productId"`
	ClosePrice decimal.Decimal `json:"closePrice"`
	// Bid        decimal.Decimal `json:"buyPrice"`
	// Ask        decimal.Decimal `json:"sellPrice"`
	// PriorSettle float64         `json:"priorSettle"`
	// Time        int64           `json:"time"`
}

type OverseaFuturesOrder struct {
	OrderNo            string    `json:"order_no"`   //cl_order_id
	Account            string    `json:"account"`    //cust_number
	ProductId          string    `json:"productId"`  //contract_code
	OrderType          string    `json:"orderType"`  //order_type
	LimitPrice         float64   `json:"limitPrice"` // order_price
	StopLossPrice      float64   `json:"stopLossPrice"`
	TakeProfitPrice    float64   `json:"takeProfitPrice"`
	InitialMargin      float64   `json:"initialMargin"`
	InitialMarginTotal float64   `json:"initialMarginTotal"`
	BuySell            string    `json:"buySell"` //side
	Leverage           int       `json:"leverage"`
	Amount             float64   `json:"amount"` //commit_qty
	Time               time.Time `json:"time"`
	Status             string    `json:"status"` //order_status
	Fee                float64   `json:"fee"`    //0
	CheckTime          string    `json:"checkTime"`
	IsDemo             int       `json:"isDemo"`
}

type CloseOrder struct {
	ClOrderID   string  `json:"clOrderID"`
	ClosePrice  float64 `json:"closePrice"`
	OpenPrice   float64 `json:"openPrice"`
	CloseAmount float64 `json:"closeAmount"`
	Status      int     `json:"status"`
}

type CancelOrder struct {
	OrderNo string `json:"orderNo"`
	Status  string `json:"status"`
}

type SystemOrder struct {
	Type      string  `json:"type"` //in(account|positionId|productId)
	Value     string  `json:"value"`
	BuySell   string  `json:"buySell"`
	Price     float64 `json:"price"`
	ProductId string  `json:"productId"`
}

type MemberPosition struct {
	ProductId  string  `json:"ProductId"`
	ClosePrice float64 `json:"ClosePrice"`
	Amount     float64 `json:"Amount"`
	Margin     float64 `json:"Margin"`
	BuySell    string  `json:"BuySell"`
	PipCost    float64 `json:"PipCost"`
	Currency   string  `json:"Currency"`
	Profit     float64 `json:"Profit"`
	Account    string  `json:"Account"`
}

type Quote struct {
	Time           time.Time `json:"time"`
	ProductId      string    `json:"productId"`
	YesterdayPrice float64   `json:"yesterdayPrice"` //昨收
	OpenPrice      float64   `json:"openPrice"`      //今開
	HighPrice      float64   `json:"highPrice"`      //最高
	LowPrice       float64   `json:"lowPrice"`       //最低
	ClosePrice     float64   `json:"closePrice"`     //成交
	Quantity       int       `json:"quantity"`       //成交量
	BuyPrice       float64   `json:"buyPrice"`       //買價
	BuyQuantity    int       `json:"buyQuantity"`
	SellPrice      float64   `json:"sellPrice"` //賣價
	SellQuantity   int       `json:"sellQuantity"`
	Volume         int       `json:"volume"`
	PriorSettle    float64   `json:"priorSettle"` // 最後交易結算
}

func (order *OverseaFuturesOrder) Insert(queue *[]*OverseaFuturesOrder) bool {
	switch order.OrderType {
	case "market":
		*queue = append(*queue, order)
	case "stop":
		*queue = append(*queue, order)
	case "limit":
		var insertIndex int
		insertIndex = -1
		if order.BuySell == "buy" {
			//   299  298 大於改小於
			for i, queuedOrder := range *queue {
				if order.LimitPrice > queuedOrder.LimitPrice {
					insertIndex = i
					break
				} else if order.LimitPrice == queuedOrder.LimitPrice {
					if order.Time.Before(queuedOrder.Time) {
						insertIndex = i
						break
					}
				}
			}
		} else {
			for i, queuedOrder := range *queue {
				if order.LimitPrice < queuedOrder.LimitPrice {
					insertIndex = i
					break
				} else if order.LimitPrice == queuedOrder.LimitPrice {
					if order.Time.Before(queuedOrder.Time) {
						insertIndex = i
						break
					}
				}
			}
		}

		if insertIndex == -1 {
			*queue = append(*queue, order)
		} else if insertIndex == 0 {
			var queue2 []*OverseaFuturesOrder
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

func (order *OverseaFuturesOrder) Remove(queue *[]*OverseaFuturesOrder) bool {
	var removeIndex int
	var removeStatus bool

	for i, queuedOrder := range *queue {
		// check order identity
		if order.OrderNo == queuedOrder.OrderNo {
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
