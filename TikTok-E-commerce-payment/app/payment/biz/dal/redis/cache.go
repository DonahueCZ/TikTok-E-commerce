package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// SetCache 存储数据到 Redis
func (r *RedisClient) SetCache(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := r.Client.Set(ctx, key, value, expiration).Err()
	return err
}

// GetCache 从 Redis 获取数据
func (r *RedisClient) GetCache(ctx context.Context, key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // 如果没有找到，返回空值
	}
	return val, err
}

// DeleteCache 删除缓存
func (r *RedisClient) DeleteCache(ctx context.Context, key string) error {
	err := r.Client.Del(ctx, key).Err()
	return err
}
