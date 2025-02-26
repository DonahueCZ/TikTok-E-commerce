package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

// RedisClient 用于封装 Redis 客户端实例
type RedisClient struct {
	Client *redis.Client
}

// NewRedisClient 创建 Redis 客户端实例
func NewRedisClient(redisAddr, redisPassword string, redisDB int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword, // No password set
		DB:       redisDB,       // Default DB
	})

	// 测试连接
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	return &RedisClient{Client: client}
}
