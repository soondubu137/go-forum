package redis

import (
	"fmt"

	"github.com/SoonDubu923/go-forum/config"

	"github.com/go-redis/redis"
)

var rds *redis.Client

func Init(cfg *config.RedisConfig) (err error) {
    rds = redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
        Password: cfg.Password,
        DB:       cfg.DB,
        PoolSize: cfg.PoolSize,
    })
    _, err = rds.Ping().Result()
    return
}

func Close() {
    _ = rds.Close()
}