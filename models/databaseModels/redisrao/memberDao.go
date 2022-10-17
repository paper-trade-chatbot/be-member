package redisrao

import (
	"context"

	"github.com/paper-trade-chatbot/be-member/cache"
	"github.com/paper-trade-chatbot/be-member/logging"
	"github.com/shopspring/decimal"
)

func GetMemberDirAgent(r *cache.RedisInstance, memberAccount string) string {
	return r.HGet(context.Background(), "Member:"+memberAccount, "DirectAgent").Val()
}

func GetMemberWarningLevel(r *cache.RedisInstance, memberAccount string) decimal.Decimal {
	result := r.HGet(context.Background(), "Member:"+memberAccount, "WarningLevel").Val()
	rDec, err := decimal.NewFromString(result)
	if err != nil {
		logging.Error("GetMemberWarningLevel error, result: %+v", result)
		rDec = decimal.Zero
	}
	return rDec
}

func GetMemberStopLossLevel(r *cache.RedisInstance, memberAccount string) decimal.Decimal {
	result := r.HGet(context.Background(), "Member:"+memberAccount, "StopLossLevel").Val()
	rDec, err := decimal.NewFromString(result)
	if err != nil {
		logging.Error("GetMemberStopLossLevel error, result: %+v", result)
		rDec = decimal.Zero
	}
	return rDec
}

func GetMarketInterval(r *cache.RedisInstance) string {
	result := r.HGet(context.Background(), "OrderSecond:future", "MarketOrderInterval").Val()
	if result == "" {
		return "0s"
	} else {
		return result + "s"
	}
}

func GetLimitInterval(r *cache.RedisInstance) string {
	result := r.HGet(context.Background(), "OrderSecond:future", "LimitOrderInterval").Val()
	if result == "" {
		return "0s"
	} else {
		return result + "s"
	}
}

func GetMemberRealizeProfit(r *cache.RedisInstance, memberAccount string) string {
	result1 := r.HGet(context.Background(), "Member:"+memberAccount, "RealizedProfitDay").Val()
	rDec1, err := decimal.NewFromString(result1)
	if err != nil {
		logging.Error("GetMemberRealizeProfit error, result1: %+v", result1)
		rDec1 = decimal.Zero
	}
	result2 := r.HGet(context.Background(), "Member:"+memberAccount, "RealizedProfitNight").Val()
	rDec2, err := decimal.NewFromString(result2)
	if err != nil {
		logging.Error("GetMemberRealizeProfit error, result2: %+v", result2)
		rDec2 = decimal.Zero
	}
	return rDec1.Add(rDec2).String()
}

func GetMemberDone(r *cache.RedisInstance, memberAccount string) string {
	result := r.HGet(context.Background(), "Member:"+memberAccount, "Done").Val()
	return result
}

func GetMemberProfit(r *cache.RedisInstance, memberAccount string) decimal.Decimal {
	result := r.HGet(context.Background(), "Member:"+memberAccount, "Profit").Val()
	rDec, err := decimal.NewFromString(result)
	if err != nil {
		logging.Error("GetMemberProfit error, result: %+v", result)
		rDec = decimal.Zero
	}
	return rDec
}

func ModifyMemberProfit(r *cache.RedisInstance, memberAccount string, profit string) error {
	err := r.HSet(context.Background(), "Member:"+memberAccount, "Profit", profit).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetSystemOrderLock(r *cache.RedisInstance, memberAccount string) string {
	result := r.HGet(context.Background(), "Member:"+memberAccount, "SystemOrderLock").Val()
	return result
}

func GetMemberFee(r *cache.RedisInstance, memberAccount string) decimal.Decimal {
	result1 := r.HGet(context.Background(), "Member:"+memberAccount, "FeeDay").Val()
	rDec1, err := decimal.NewFromString(result1)
	if err != nil {
		logging.Error("GetMemberFee error, result1: %+v", result1)
		rDec1 = decimal.Zero
	}
	result2 := r.HGet(context.Background(), "Member:"+memberAccount, "FeeNight").Val()
	rDec2, err := decimal.NewFromString(result2)
	if err != nil {
		logging.Error("GetMemberFee error, result2: %+v", result2)
		rDec2 = decimal.Zero
	}
	return rDec1.Add(rDec2)
}

func GetMemberWalletAmount(r *cache.RedisInstance, memberAccount string) decimal.Decimal {
	result := r.HGet(context.Background(), "Member:"+memberAccount, "WalletAmount").Val()
	rDec, err := decimal.NewFromString(result)
	if err != nil {
		logging.Error("GetMemberWalletAmount error, result: %+v", result)
		rDec = decimal.Zero
	}
	return rDec
}

func GetMemberPrepayFee(r *cache.RedisInstance, memberAccount string) decimal.Decimal {
	result := r.HGet(context.Background(), "Member:"+memberAccount, "PrepayFee").Val()
	rDec, err := decimal.NewFromString(result)
	if err != nil {
		logging.Error("GetMemberPrepayFee error, result: %+v", result)
		rDec = decimal.Zero
	}
	return rDec
}

func GetMemberMargin(r *cache.RedisInstance, memberAccount string) decimal.Decimal {
	result := r.HGet(context.Background(), "Member:"+memberAccount, "Margin").Val()
	rDec, err := decimal.NewFromString(result)
	if err != nil {
		logging.Error("GetMemberMargin error, result: %+v", result)
		rDec = decimal.Zero
	}
	return rDec
}

func GetMemberTotalAmount(r *cache.RedisInstance, memberAccount string) decimal.Decimal {
	result := r.HGet(context.Background(), "Member:"+memberAccount, "TotalAmount").Val()
	rDec, err := decimal.NewFromString(result)
	if err != nil {
		logging.Error("GetMemberTotalAmount error, result: %+v", result)
		rDec = decimal.Zero
	}
	return rDec
}

func GetMemberPrincipal(r *cache.RedisInstance, memberAccount string) decimal.Decimal {
	result := r.HGet(context.Background(), "Member:"+memberAccount, "Principal").Val()
	rDec, err := decimal.NewFromString(result)
	if err != nil {
		logging.Error("GetMemberPrincipal error, result: %+v", result)
		rDec = decimal.Zero
	}
	return rDec
}

func ModifyRiskWarningNotifyLock(r *cache.RedisInstance, memberAccount string, status string) error {
	err := r.Set(context.Background(), "Member:"+memberAccount+":RiskWarningNotifyLock", status, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func CheckExistRiskWarningNotifyLock(r *cache.RedisInstance, memberAccount string) bool {
	exist := r.Exists(context.Background(), "Member:"+memberAccount+":RiskWarningNotifyLock").Val()

	if exist == 1 {
		return true
	}
	return false
}

func DeleteRiskWarningNotifyLock(r *cache.RedisInstance, memberAccount string) {
	err := r.Del(context.Background(), "Member:"+memberAccount+":RiskWarningNotifyLock").Err()
	if err != nil {
		logging.Error("err %v", err)
	}
	//logging.Info("Member:" + memberAccount + ":RiskWarningNotifyLock")
}

func ModifyRiskStopLossNotifyLock(r *cache.RedisInstance, memberAccount string, status string) error {
	err := r.Set(context.Background(), "Member:"+memberAccount+":RiskStopLossNotifyLock", status, -1).Err()
	if err != nil {
		return err
	}
	return nil
}

func CheckExistRiskStopLossNotifyLock(r *cache.RedisInstance, memberAccount string) bool {
	exist := r.Exists(context.Background(), "Member:"+memberAccount+":RiskStopLossNotifyLock").Val()
	if exist == 1 {
		return true
	}
	return false
}

func DeleteRiskStopLossNotifyLock(r *cache.RedisInstance, memberAccount string) {
	err := r.Del(context.Background(), "Member:"+memberAccount+":RiskStopLossNotifyLock").Err()
	if err != nil {
		logging.Error("err %v", err)
	}
	//logging.Info("Member:" + memberAccount + ":RiskStopLossNotifyLock")
}
