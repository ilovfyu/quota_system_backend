package common

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	g "quota_system/global"
)

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", g.ServiceConfig.RdsConfig.Host, g.ServiceConfig.RdsConfig.Port),
		DB:   g.ServiceConfig.RdsConfig.DB,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		zap.L().Error("rdb ping failed", zap.Error(err))
		panic(err)
	}
	g.RDS = rdb
}
