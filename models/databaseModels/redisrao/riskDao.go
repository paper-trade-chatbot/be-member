package redisrao

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/paper-trade-chatbot/be-member/cache"
	"github.com/paper-trade-chatbot/be-member/service/utils/jsonformat"
	"github.com/spf13/cast"
)

func GetMembersByProduct(r *cache.RedisInstance, productId string) []string {
	result := r.SMembers(context.Background(), "risk-product:"+productId).Val()
	return result
}

func ModifyPositionProfit(r *cache.RedisInstance, account string, productId string, profit float64) error {
	err := r.HSet(context.Background(), "risk-member:"+account+":"+productId, "Profit", profit).Err()
	if err != nil {
		return err
	}
	return nil
}

func ModifyPosition(r *cache.RedisInstance, position *jsonformat.MemberPosition) error {
	err := r.HMSet(context.Background(), "risk-member:"+position.Account+":"+position.ProductId, map[string]interface{}{
		"ClosePrice": position.ClosePrice,
		"Amount":     position.Amount,
		"Margin":     position.Margin,
	}).Err()
	if err != nil {
		return err
	}
	return nil
}

func SetMemberProfitRank(r *cache.RedisInstance, account string, profit string) error {
	err := r.ZAdd(context.Background(), "Profit-Rank", &redis.Z{
		Score:  cast.ToFloat64(profit),
		Member: account,
	}).Err()

	if err != nil {
		return err
	}
	return nil
}

func RemMemberProfitRank(r *cache.RedisInstance, account string) error {
	err := r.ZRem(context.Background(), "Profit-Rank", account).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetMemberProfitRank(r *cache.RedisInstance, order string) []string {
	var result []string

	if order == "loss" {
		result = r.ZRange(context.Background(), "Profit-Rank", 0, 99).Val()
	} else {
		result = r.ZRevRange(context.Background(), "Profit-Rank", 0, 99).Val()
	}
	return result
}

func SetMemberProfitPercentRank(r *cache.RedisInstance, account string, profit string) error {
	err := r.ZAdd(context.Background(), "ProfitPercent-Rank", &redis.Z{
		Score:  cast.ToFloat64(profit),
		Member: account,
	}).Err()

	if err != nil {
		return err
	}
	return nil
}

func RemMemberProfitPercentRank(r *cache.RedisInstance, account string) error {
	err := r.ZRem(context.Background(), "ProfitPercent-Rank", account).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetMemberProfitPercentRank(r *cache.RedisInstance, order string) []string {
	var result []string

	if order == "loss" {
		result = r.ZRange(context.Background(), "ProfitPercent-Rank", 0, 100).Val()
	} else {
		result = r.ZRevRange(context.Background(), "ProfitPercent-Rank", 0, 100).Val()
	}
	return result
}

func SetProductLock(r *cache.RedisInstance, productId string) error {
	err := r.SAdd(context.Background(), "Product-Lock", productId).Err()
	if err != nil {
		return err
	}
	return nil
}

func DelProductLock(r *cache.RedisInstance, productId string) error {
	err := r.SRem(context.Background(), "Product-Lock", productId).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetProductLock(r *cache.RedisInstance, productId string) bool {
	return r.SIsMember(context.Background(), "Product-Lock", productId).Val()
}
