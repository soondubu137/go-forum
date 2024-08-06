package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rds *redis.Client

func Init() (err error) {
    rds = redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
        Password: viper.GetString("redis.password"),
        DB:       viper.GetInt("redis.db"),
        PoolSize: viper.GetInt("redis.pool_size"),
    })
    _, err = rds.Ping().Result()
    return
}

func Close() {
    _ = rds.Close()
}