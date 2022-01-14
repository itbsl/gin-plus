package redis

import (
	"context"
	"gin-plus/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"strconv"
)

var rdb *redis.Client

func Init(cfg *config.RedisConfig) (err error) {
	ctx := context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		zap.L().Error("Connect redis failed.", zap.Error(err))
	}
	return
}

func Close() {
	_ = rdb.Close()
}
