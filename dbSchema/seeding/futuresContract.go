package main

import (
	"context"

	"github.com/paper-trade-chatbot/be-member/cache"
	"github.com/paper-trade-chatbot/be-member/database"
	"github.com/paper-trade-chatbot/be-member/logging"
	models "github.com/paper-trade-chatbot/be-member/models/databaseModels"
	"github.com/paper-trade-chatbot/be-member/service/gormdao/futurescontractdao"
)

func futuresContractSeeder() {
	tx := database.GetDB().Begin()
	r, _ := cache.GetRedis()

	defer func() {
		tx.RollbackUnlessCommitted()
		if r := recover(); r != nil {
			logging.Debug("%s", r)
		}
	}()

	data := map[string]string{
		"HSIA20":    "HKFX",
		"HMHA20":    "HKFX",
		"HIHHIFA20": "HKFX",
		"HIMCHFA20": "HKFX",
		"CNA20":     "SGX",
		"NQA20":     "CME",
		"YMA20":     "CME",
		"SPA20":     "CME",
		"ESA20":     "CME",
		"N225A20":   "JPX",
		"DAXA20":    "EUREX",
		"GCA20":     "NYMEX",
		"SIA20":     "NYMEX",
		"ASDA20":    "NYMEX",
		"FSFA20":    "NYMEX",
		"FITXA20":   "TAIFEX",
		"FIMTXA20":  "TAIFEX",
	}

	data2 := map[string]string{
		"HSIA20":    "index",
		"HMHA20":    "index",
		"HIHHIFA20": "index",
		"HIMCHFA20": "index",
		"CNA20":     "index",
		"NQA20":     "index",
		"YMA20":     "index",
		"SPA20":     "index",
		"ESA20":     "index",
		"N225A20":   "index",
		"DAXA20":    "index",
		"GCA20":     "metal",
		"SIA20":     "metal",
		"ASDA20":    "metal",
		"FSFA20":    "metal",
		"FITXA20":   "index",
		"FIMTXA20":  "index",
	}

	for contract, exahange := range data {

		if futurescontractdao.Get(tx, &futurescontractdao.QueryModel{
			ContractCode: contract,
		}) == nil {

			futurescontractdao.New(tx, &models.FuturesContractModel{
				ExchangeCode:      exahange,
				ContractCode:      contract,
				ContractStatus:    "enabled",
				FutureVarietyCode: data2[contract],
			})

			r.HSet(context.Background(), "contract", contract, exahange)

		}
	}
	tx.Commit()

}
