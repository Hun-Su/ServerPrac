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
func GetValue(cli *redis.Client, key string) *redis.StringCmd {
	val := cli.Get(cli.Context(), key)
	return val
}

//leehs 20220530 Redis 데이터 저장
func SetValue(cli *redis.Client, key string, val interface{}) string {
	t, _ := time.ParseDuration(CONFIG.Redis.Timeout)

	status := cli.Set(cli.Context(), key, val, t)
	return status.Val()
}

//leehs 20220530 Redis 데이터 삭제
func Empty(cli *redis.Client) {
	cli.FlushDBAsync(cli.Context())
}
