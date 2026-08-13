package cache

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheClient interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	IsAvailable() bool
}

type redisCache struct {
	client    *redis.Client
	available bool
}

func NewRedisCache() CacheClient {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	redisPassword := os.Getenv("REDIS_PASSWORD")

	rdb := redis.NewClient(&redis.Options{
		Addr:         redisAddr,
		Password:     redisPassword,
		DB:           0,
		MaxRetries:   0,
		DialTimeout:  200 * time.Millisecond,
		ReadTimeout:  200 * time.Millisecond,
		WriteTimeout: 200 * time.Millisecond,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	available := true
	if err != nil {
		log.Printf("⚠️ [Redis Cache Notice]: Không thể kết nối Redis (%s). Hệ thống sẽ tự động dùng Fallback Pass-Through Cache.", redisAddr)
		available = false
	} else {
		log.Printf("⚡ [Redis Cache]: Kết nối Redis thành công tại %s!", redisAddr)
	}

	return &redisCache{
		client:    rdb,
		available: available,
	}
}

func (r *redisCache) IsAvailable() bool {
	return r.available
}

func (r *redisCache) Get(ctx context.Context, key string) (string, error) {
	if !r.available {
		return "", fmt.Errorf("redis offline")
	}
	return r.client.Get(ctx, key).Result()
}

func (r *redisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	if !r.available {
		return nil
	}
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *redisCache) Delete(ctx context.Context, key string) error {
	if !r.available {
		return nil
	}
	return r.client.Del(ctx, key).Err()
}
