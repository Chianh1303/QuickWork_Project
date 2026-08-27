package cache

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
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
	redisURL := strings.Trim(os.Getenv("REDIS_URL"), "\"'\t\n\r ")
	redisURL = strings.TrimPrefix(redisURL, "REDIS_URL=")
	redisURL = strings.Trim(redisURL, "\"'\t\n\r ")

	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	var rdb *redis.Client

	if redisURL != "" {
		opt, err := redis.ParseURL(redisURL)
		if err == nil {
			opt.DialTimeout = 3 * time.Second
			opt.ReadTimeout = 3 * time.Second
			opt.WriteTimeout = 3 * time.Second
			rdb = redis.NewClient(opt)
			redisAddr = opt.Addr
		} else {
			log.Printf("⚠️ [Redis URL Parse Error]: %v", err)
		}
	}

	if rdb == nil {
		if redisAddr == "" {
			redisAddr = "localhost:6379"
		}
		rdb = redis.NewClient(&redis.Options{
			Addr:         redisAddr,
			Password:     redisPassword,
			DB:           0,
			MaxRetries:   0,
			DialTimeout:  1 * time.Second,
			ReadTimeout:  1 * time.Second,
			WriteTimeout: 1 * time.Second,
		})
	}

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
