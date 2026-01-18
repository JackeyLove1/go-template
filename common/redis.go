package common

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/songquanpeng/one-api/common/logger"
)

var RDB redis.Cmdable
var RedisEnabled = true

// InitRedisClient This function is called after init()
func InitRedisClient() (err error) {
	if os.Getenv("REDIS_CONN_STRING") == "" {
		RedisEnabled = false
		logger.SysLog("REDIS_CONN_STRING not set, Redis is not enabled")
		return nil
	}
	if os.Getenv("SYNC_FREQUENCY") == "" {
		RedisEnabled = false
		logger.SysLog("SYNC_FREQUENCY not set, Redis is disabled")
		return nil
	}
	redisConnString := os.Getenv("REDIS_CONN_STRING")
	if os.Getenv("REDIS_MASTER_NAME") == "" {
		logger.SysLog("Redis is enabled")
		opt, err := redis.ParseURL(redisConnString)
		if err != nil {
			logger.FatalLog("failed to parse Redis connection string: " + err.Error())
		}
		RDB = redis.NewClient(opt)
	} else {
		// cluster mode
		logger.SysLog("Redis cluster mode enabled")
		RDB = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:      strings.Split(redisConnString, ","),
			Password:   os.Getenv("REDIS_PASSWORD"),
			MasterName: os.Getenv("REDIS_MASTER_NAME"),
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		logger.FatalLog("Redis ping test failed: " + err.Error())
	}
	return err
}

func ParseRedisOption() *redis.Options {
	opt, err := redis.ParseURL(os.Getenv("REDIS_CONN_STRING"))
	if err != nil {
		logger.FatalLog("failed to parse Redis connection string: " + err.Error())
	}
	return opt
}

func RedisSet(key string, value string, expiration time.Duration) error {
	ctx := context.Background()
	return RDB.Set(ctx, key, value, expiration).Err()
}

func RedisGet(key string) (string, error) {
	ctx := context.Background()
	return RDB.Get(ctx, key).Result()
}

func RedisDel(key string) error {
	ctx := context.Background()
	return RDB.Del(ctx, key).Err()
}

func RedisDecrease(key string, value int64) error {
	ctx := context.Background()
	return RDB.DecrBy(ctx, key, value).Err()
}

// ----------------------------
// Hash (Redis Hash Table)
// ----------------------------

func RedisHSet(key string, field string, value string) error {
	ctx := context.Background()
	return RDB.HSet(ctx, key, field, value).Err()
}

func RedisHGet(key string, field string) (string, error) {
	ctx := context.Background()
	return RDB.HGet(ctx, key, field).Result()
}

func RedisHGetAll(key string) (map[string]string, error) {
	ctx := context.Background()
	return RDB.HGetAll(ctx, key).Result()
}

func RedisHDel(key string, fields ...string) error {
	ctx := context.Background()
	return RDB.HDel(ctx, key, fields...).Err()
}

// ----------------------------
// Set (Redis Set)
// ----------------------------

func RedisSAdd(key string, members ...string) error {
	ctx := context.Background()
	args := make([]interface{}, 0, len(members))
	for _, m := range members {
		args = append(args, m)
	}
	return RDB.SAdd(ctx, key, args...).Err()
}

func RedisSMembers(key string) ([]string, error) {
	ctx := context.Background()
	return RDB.SMembers(ctx, key).Result()
}

func RedisSRem(key string, members ...string) error {
	ctx := context.Background()
	args := make([]interface{}, 0, len(members))
	for _, m := range members {
		args = append(args, m)
	}
	return RDB.SRem(ctx, key, args...).Err()
}

func RedisSIsMember(key string, member string) (bool, error) {
	ctx := context.Background()
	return RDB.SIsMember(ctx, key, member).Result()
}

// ----------------------------
// ZSet (Redis Sorted Set)
// ----------------------------

func RedisZAdd(key string, score float64, member string) error {
	ctx := context.Background()
	return RDB.ZAdd(ctx, key, &redis.Z{Score: score, Member: member}).Err()
}

func RedisZRange(key string, start int64, stop int64) ([]string, error) {
	ctx := context.Background()
	return RDB.ZRange(ctx, key, start, stop).Result()
}

func RedisZRevRange(key string, start int64, stop int64) ([]string, error) {
	ctx := context.Background()
	return RDB.ZRevRange(ctx, key, start, stop).Result()
}

func RedisZRem(key string, members ...string) error {
	ctx := context.Background()
	args := make([]interface{}, 0, len(members))
	for _, m := range members {
		args = append(args, m)
	}
	return RDB.ZRem(ctx, key, args...).Err()
}

func RedisZScore(key string, member string) (float64, error) {
	ctx := context.Background()
	return RDB.ZScore(ctx, key, member).Result()
}

// ----------------------------
// String / Key
// ----------------------------

func RedisSetNX(key string, value string, expiration time.Duration) (bool, error) {
	ctx := context.Background()
	return RDB.SetNX(ctx, key, value, expiration).Result()
}

func RedisExpire(key string, expiration time.Duration) (bool, error) {
	ctx := context.Background()
	return RDB.Expire(ctx, key, expiration).Result()
}

func RedisTTL(key string) (time.Duration, error) {
	ctx := context.Background()
	return RDB.TTL(ctx, key).Result()
}

func RedisIncr(key string) (int64, error) {
	ctx := context.Background()
	return RDB.Incr(ctx, key).Result()
}

func RedisIncrBy(key string, increment int64) (int64, error) {
	ctx := context.Background()
	return RDB.IncrBy(ctx, key, increment).Result()
}

// ----------------------------
// List
// ----------------------------

func RedisLPush(key string, values ...string) (int64, error) {
	ctx := context.Background()
	args := make([]interface{}, 0, len(values))
	for _, v := range values {
		args = append(args, v)
	}
	return RDB.LPush(ctx, key, args...).Result()
}

func RedisRPush(key string, values ...string) (int64, error) {
	ctx := context.Background()
	args := make([]interface{}, 0, len(values))
	for _, v := range values {
		args = append(args, v)
	}
	return RDB.RPush(ctx, key, args...).Result()
}

func RedisLPop(key string) (string, error) {
	ctx := context.Background()
	return RDB.LPop(ctx, key).Result()
}

func RedisRPop(key string) (string, error) {
	ctx := context.Background()
	return RDB.RPop(ctx, key).Result()
}

func RedisLLen(key string) (int64, error) {
	ctx := context.Background()
	return RDB.LLen(ctx, key).Result()
}

func RedisLRange(key string, start int64, stop int64) ([]string, error) {
	ctx := context.Background()
	return RDB.LRange(ctx, key, start, stop).Result()
}

// ----------------------------
// Bitmap
// ----------------------------

func RedisSetBit(key string, offset int64, value int) (int64, error) {
	ctx := context.Background()
	return RDB.SetBit(ctx, key, offset, value).Result()
}

func RedisGetBit(key string, offset int64) (int64, error) {
	ctx := context.Background()
	return RDB.GetBit(ctx, key, offset).Result()
}

// RedisBitCount counts set bits.
// If you want the whole string, pass start=0 and end=-1.
func RedisBitCount(key string, start int64, end int64) (int64, error) {
	ctx := context.Background()
	return RDB.BitCount(ctx, key, &redis.BitCount{Start: start, End: end}).Result()
}
