package initialize

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"goBlog/lib/config/model"
	"goBlog/lib/global"
)

// initRedisClient 初始化redis
func initRedisClient(redisCfg model.Redis) (redis.UniversalClient, error) {
	var client redis.UniversalClient

	// 使用单例模式
	client = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Redis 连接失败:", pong, err)
		return nil, err
	}

	return client, nil
}

// Redis 初始化redis
func Redis() redis.UniversalClient {
	redisClient, err := initRedisClient(global.GVA_CONFIG.Redis)
	if err != nil {
		panic(err)
	}
	return redisClient
}
