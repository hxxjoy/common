// pkg/cache/redis.go
package cache

import (
    "github.com/redis/go-redis/v9"
)

// RedisConfig Redis配置
type RedisConfig struct {
    Addr     string
    Password string
    DB       int
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg RedisConfig) *redis.Client {
    return redis.NewClient(&redis.Options{
        Addr:     cfg.Addr,
        Password: cfg.Password,
        DB:       cfg.DB,
    })
}