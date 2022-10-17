package main

import (
	"context"

	"github.com/paper-trade-chatbot/be-member/cache"
	"github.com/paper-trade-chatbot/be-member/database"
	"github.com/paper-trade-chatbot/be-member/logging"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
	"github.com/paper-trade-chatbot/be-member/service/gormdao/exchangedao"
	"github.com/spf13/cast"
)

func exchangeSeeder() {
	tx := database.GetDB().Begin()
	r, _ := cache.GetRedis()
	defer func() {
		tx.RollbackUnlessCommitted()
		if r := recover(); r != nil {
			logging.Debug("%s", r)
		}
	}()

	data := []map[string]interface{}{
		{
			"Symbol":        "SGX",
			"Name":          "新加坡交易所",
			"Timezone":      8,
			"OpenTime":      "0900",
			"Location":      "Asia/Singapore",
			"Type":          "futures",
			"Status":        "enabled",
			"NormalTime":    `[ { "startDay":"1", "endDay":"5", "startTime":"0900", "endTime":"1635"} ]`,
			"ExceptionTime": `{"trade":[],"stopTrade":[]}`,
		},

		{
			"Symbol":        "CME",
			"Name":          "芝加哥商品交易所",
			"Timezone":      -6,
			"OpenTime":      "0930",
			"Location":      "America/Chicago",
			"Type":          "futures",
			"Status":        "enabled",
			"NormalTime":    `[ { "startDay":"1", "endDay":"5", "startTime":"0700", "endTime":"0515"}, { "startDay":"1", "endDay":"5", "startTime":"0530", "endTime":"0600"} ]`,
			"ExceptionTime": `{"trade":[],"stopTrade":[]}`,
		},
		{
			"Symbol":        "NYMEX",
			"Name":          "紐約商業交易所",
			"Timezone":      -5,
			"OpenTime":      "0930",
			"Location":      "America/New_York",
			"Type":          "futures",
			"Status":        "enabled",
			"NormalTime":    `[ { "startDay":"1", "endDay":"5", "startTime":"0700", "endTime":"0600" } ]`,
			"ExceptionTime": `{"trade":[],"stopTrade":[]}`,
		},
		{
			"Symbol":        "HKFX", //HKEx
			"Name":          "香港期貨交易所",
			"Timezone":      8,
			"OpenTime":      "0930",
			"Location":      "Asia/Hong_Kong",
			"Type":          "futures",
			"Status":        "enabled",
			"NormalTime":    `[ { "startDay":"1", "endDay":"5", "startTime":"0915", "endTime":"1200" },{ "startDay":"1", "endDay":"5", "startTime":"1300", "endTime":"1630" }]`,
			"ExceptionTime": `{"trade":[],"stopTrade":[]}`,
		},
		{
			"Symbol":        "EUREX",
			"Name":          "歐洲期貨交易所",
			"Timezone":      1,
			"OpenTime":      "0930",
			"Location":      "Europe/Berlin",
			"Type":          "futures",
			"Status":        "enabled",
			"NormalTime":    `[ { "startDay":"1", "endDay":"5", "startTime":"0815", "endTime":"0400" } ]`,
			"ExceptionTime": `{"trade":[],"stopTrade":[]}`,
		},
		{
			"Symbol":        "LME",
			"Name":          "伦敦金属交易所",
			"Timezone":      -5,
			"OpenTime":      "0930",
			"Location":      "Europe/London",
			"Type":          "futures",
			"Status":        "enabled",
			"NormalTime":    `[ { "startDay":"1", "endDay":"5", "startTime":"0700", "endTime":"0400" } ]`,
			"ExceptionTime": `{"trade":[],"stopTrade":[]}`,
		},
		{
			"Symbol":        "TAIFEX",
			"Name":          "台湾期货交易所",
			"Timezone":      8,
			"OpenTime":      "0930",
			"Location":      "America/New_York",
			"Type":          "futures",
			"Status":        "enabled",
			"NormalTime":    `[ { "startDay":"1", "endDay":"5", "startTime":"0700", "endTime":"0400" } ]`,
			"ExceptionTime": `{"trade":[],"stopTrade":[]}`,
		},
		{
			"Symbol":        "JPX",
			"Name":          "日本期货交易所",
			"Timezone":      9,
			"OpenTime":      "0930",
			"Location":      "Asia/Tokyo",
			"Type":          "futures",
			"Status":        "enabled",
			"NormalTime":    `[ { "startDay":"1", "endDay":"5", "startTime":"0700", "endTime":"0400" } ]`,
			"ExceptionTime": `{"trade":[],"stopTrade":[]}`,
		},
	}

	for _, value := range data {
		if exchangedao.Get(tx, &exchangedao.QueryModel{
			ExchangeCode: cast.ToString(value["Symbol"]),
		}) == nil {
			exchange := &models.ExchangeModel{
				ExchangeCode:   cast.ToString(value["Symbol"]),
				ExchangeStatus: cast.ToString(value["Status"]),
				TradeTime:      cast.ToString(value["NormalTime"]),
				ExceptionTime:  cast.ToString(value["ExceptionTime"]),
				ExchangeName:   cast.ToString(value["Name"]),
				Timezone:       cast.ToFloat64(value["Timezone"]),
				OpenTime:       cast.ToString(value["OpenTime"]),
				Location:       cast.ToString(value["Location"]),
			}
			exchangedao.New(tx, exchange)
			r.HMSet(context.Background(), "exchange:"+exchange.ExchangeCode, map[string]interface{}{
				"timeZone": exchange.Timezone,
				"openTime": exchange.OpenTime,
			})

		}

	}

	tx.Commit()
}
