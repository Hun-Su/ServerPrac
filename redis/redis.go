package redis

import (
	"echo/Config"
	"github.com/go-redis/redis/v8"
	"time"
)

var CONFIG = Config.LoadConfig()

//leehs 20220530 Redis 클라이언트 정보 가져오기
func GetRedisCli() *redis.Client {
	t, _ := time.ParseDuration(CONFIG.Redis.Timeout)
	client := redis.NewClient(&redis.Options{
		Addr:        CONFIG.Redis.Port,
		ReadTimeout: time.Second * t,
	})
	return client
}

//leehs 20220530 Redis 데이터 가져오기
func GetValue(cli *redis.Client, key string) string {
	ctx := cli.Context()

	val := cli.Get(ctx, key)
	return val.Val()
}

//leehs 20220530 Redis 데이터 저장
func SetValue(cli *redis.Client, key string, val interface{}) string {
	ctx := cli.Context()
	t, _ := time.ParseDuration(CONFIG.Redis.Timeout)

	status := cli.Set(ctx, key, val, t)
	return status.Val()
}

// Set
// Get
